package main

import (
	"fmt"
	"sync"
)

func producer9() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func consumer9(size int, c <-chan int) <-chan int {
	out := make(chan int)

	wg := sync.WaitGroup{}

	work := func() {
		for n := range c {
			out <- n * n
		}

		wg.Done()
	}

	wg.Add(size)
	for i := 0; i < size; i++ {
		go work()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	c := producer()
	result := consumer(5, c)

	for v := range result {
		fmt.Println(v)
	}
}
