package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(slice))
	for _, i := range slice {
		go func(j int) {
			defer wg.Done()
			fmt.Println(j * j)
		}(i)
	}

	wg.Wait()
}
