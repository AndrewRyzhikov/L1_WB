package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func producer4(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				ch <- 2
				time.Sleep(time.Second)
			}
		}
	}()

	return ch
}

func consumer4(wg *sync.WaitGroup, len int, ctx context.Context, ch <-chan int) {
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
	n := 2
	ctx, cancel := context.WithCancel(context.Background())

	ch := producer4(ctx)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	wg := &sync.WaitGroup{}
	consumer4(wg, n, ctx, ch)

	<-shutdown
	cancel()

	wg.Wait()
}
