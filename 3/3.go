package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(slice))

	var sum int
	m := sync.Mutex{}

	for _, i := range slice {
		go func() {
			defer wg.Done()
			m.Lock()
			sum += i
			m.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}
