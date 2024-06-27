package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count uint64
}

func (c *Counter) Increment() {
	atomic.AddUint64(&c.count, 1)
}

func (c *Counter) Count() uint64 {
	return atomic.LoadUint64(&c.count)
}

func main() {
	var n int = 100

	var wg sync.WaitGroup
	counter := Counter{count: 0}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			for c := 0; c < 100; c++ {
				counter.Increment()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter.Count())
}
