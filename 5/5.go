package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func producer5(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- 2
			}
		}
	}()

	return ch
}

func consumer5(wg *sync.WaitGroup, len int, ctx context.Context, ch <-chan int) {
	wg.Add(len)

	work := func(numberWorker int) {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case x, ok := <-ch:
				if !ok {
					return
				}
				fmt.Printf("worker%d: %d\n", numberWorker, x)
			}
		}
	}

	for i := 0; i < len; i++ {
		go work(i)
	}
}

func main() {
	n := 1
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second)
	defer cancel()

	ch := producer5(ctx)

	wg := &sync.WaitGroup{}
	consumer5(wg, n, ctx, ch)

	wg.Wait()
}
