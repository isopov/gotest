package main

import (
	"strconv"
	"sync/atomic"
	"time"
)

const submiters = 12
const submitterWait = time.Millisecond

type batcher interface {
	Submit(val string) string
}

func main() {
	for i := 1; i <= 8; i++ {
		println("batcher threads " + strconv.Itoa(i))
		batcher := newBatcher(process, i, 10, 100*time.Millisecond)

		measurement := measure(batcher)

		println(measurement.t, measurement.avgBatchSize())

		batcher = newBatcher2(process, i, 10, 100*time.Millisecond)

		measurement = measure(batcher)

		println(measurement.t, measurement.avgBatchSize())
	}

}

type measurement struct {
	t                                   time.Duration
	processedCount, processedExecutions int64
}

func (m measurement) avgBatchSize() float64 {
	return float64(m.processedCount) / float64(m.processedExecutions)
}

func measure(batcher batcher) measurement {
	processedCount = newInt64(0)
	processedExecutions = newInt64(0)

	finishedChan := make(chan bool, submiters)
	start := time.Now()
	for i := 0; i < submiters; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				result := batcher.Submit("foo")
				if "foobar" != result {
					panic("unexpected " + result)
				}
				time.Sleep(submitterWait)
			}
			finishedChan <- true
		}()
	}
	for i := 0; i < submiters; i++ {
		<-finishedChan
	}
	return measurement{
		t:                   time.Since(start),
		processedCount:      *processedCount,
		processedExecutions: *processedExecutions,
	}
}

func newInt64(v int64) *int64 {
	r := new(int64)
	*r = v
	return r
}

var processedCount = newInt64(0)
var processedExecutions = newInt64(0)

func process(vals []string) []string {
	time.Sleep(10 * time.Millisecond)

	result := make([]string, len(vals))
	for i, val := range vals {
		result[i] = val + "bar"
	}
	atomic.AddInt64(processedExecutions, 1)
	atomic.AddInt64(processedCount, int64(len(vals)))
	return result
}

type batcher1 struct {
	f           func([]string) []string
	batchSize   int
	flushPeriod time.Duration

	inChan chan *batchRequest
}

func newBatcher(f func([]string) []string, batcherThreads, batchSize int, flushPeriod time.Duration) batcher {
	b := &batcher1{
		f:           f,
		batchSize:   batchSize,
		flushPeriod: flushPeriod,
		inChan:      make(chan *batchRequest, 1000),
	}

	for i := 0; i < batcherThreads; i++ {
		go b.process()
	}
	return b
}

func (b *batcher1) process() {
	timer := time.NewTimer(b.flushPeriod)

	var batch []*batchRequest

	reset := func() {
		batch = nil
		timer = time.NewTimer(b.flushPeriod)
	}

	for {
		select {
		case req := <-b.inChan:
			batch = append(batch, req)
			if len(batch) < b.batchSize {
				continue
			} else {
				timer.Stop()
			}
		case <-timer.C:
			if len(batch) == 0 {
				reset()
				continue
			}
		}

		vals := make([]string, len(batch))
		for i, req := range batch {
			vals[i] = req.val
		}

		results := b.f(vals)

		for i, req := range batch {
			req.resultChan <- results[i]
		}
		reset()
	}
}

func (b *batcher1) Submit(val string) string {
	request := newBatchRequest(val)
	b.inChan <- request
	result := <-request.resultChan
	return result
}

func newBatchRequest(val string) *batchRequest {
	return &batchRequest{
		val:        val,
		resultChan: make(chan string, 1), // so that a writer is not blocker by a reader
	}
}

type batchRequest struct {
	val        string
	resultChan chan string
}

type batcher2 struct {
	f           func([]string) []string
	batchSize   int
	flushPeriod time.Duration

	inChan chan *batchRequest

	batchProcessChan chan []*batchRequest
}

func newBatcher2(f func([]string) []string, batcherThreads, batchSize int, flushPeriod time.Duration) batcher {
	b := &batcher2{
		f:                f,
		batchSize:        batchSize,
		flushPeriod:      flushPeriod,
		inChan:           make(chan *batchRequest, 1000),
		batchProcessChan: make(chan []*batchRequest, 1),
	}

	go b.process()

	for i := 0; i < batcherThreads; i++ {
		go b.processBatches()
	}

	return b

}

func (b *batcher2) Submit(val string) string {
	request := newBatchRequest(val)
	b.inChan <- request
	result := <-request.resultChan
	return result
}

func (b *batcher2) process() {
	timer := time.NewTimer(b.flushPeriod)

	var batch []*batchRequest

	reset := func() {
		batch = nil
		timer = time.NewTimer(b.flushPeriod)
	}

	for {
		select {
		case req := <-b.inChan:
			batch = append(batch, req)
			if len(batch) < b.batchSize {
				continue
			} else {
				timer.Stop()
			}
		case <-timer.C:
			if len(batch) == 0 {
				reset()
				continue
			}
		}

		b.batchProcessChan <- batch
		reset()
	}
}

func (b *batcher2) processBatches() {
	for {
		batch := <-b.batchProcessChan

		vals := make([]string, len(batch))
		for i, req := range batch {
			vals[i] = req.val
		}

		results := b.f(vals)

		for i, req := range batch {
			req.resultChan <- results[i]
		}
	}
}
