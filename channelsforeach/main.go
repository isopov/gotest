package main

import (
	"sync"
	"time"
)

func main() {
	ch := make(chan string, 100)

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			counter := 0
			for range ch {
				counter++
				time.Sleep(100 * time.Millisecond)
			}
			println(counter)
			wg.Done()
		}()
	}

	for i := 0; i < 100; i++ {
		ch <- "foo"
	}
	close(ch)
	wg.Wait()
	for range ch {
		println("Should not happen")
	}
}
