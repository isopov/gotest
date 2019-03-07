package main

import (
	"golang.org/x/net/context"
	"sync/atomic"
)

func main() {
	var noFirstCheckCounter, firstCheckCounter int64 = 0, 0
	for i := 0; i < 100000; i++ {
		noFirstCheckCounter += test(noFirstCheck)
		firstCheckCounter += test(withFirstCheck)
	}
	println("with first check ", firstCheckCounter)
	println("  no first check ", noFirstCheckCounter)
}

func test(chReader func(context.Context, chan int, chan bool, *int64)) int64 {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int, 100)
	endCh := make(chan bool)

	var counter int64 = 0
	go chReader(ctx, ch, endCh, &counter)

	for i := 0; i < 100; i++ {
		ch <- i
	}
	cancel()

	<-endCh

	return counter
}


func withFirstCheck(ctx context.Context, ch chan int, endCh chan bool, counter *int64) {
	for {
		select {
		case <-ctx.Done():
			close(endCh)
			return
		default:
		}

		select {
		case <-ctx.Done():
			close(endCh)
			return
		case <-ch:
			atomic.AddInt64(counter, 1)
		}
	}
}

func noFirstCheck(ctx context.Context, ch chan int, endCh chan bool, counter *int64) {
	for {
		select {
		case <-ctx.Done():
			close(endCh)
			return
		case <-ch:
			atomic.AddInt64(counter, 1)
		}
	}
}
