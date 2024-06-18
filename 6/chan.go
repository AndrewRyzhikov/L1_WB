package five

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	stopCh := make(chan struct{})

	wg.Add(1)
	go func(stopCh chan struct{}, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-stopCh:
				fmt.Println("Goroutine stopped by channel.")
				return
			default:
				fmt.Println("Working...")
				time.Sleep(1 * time.Second)
			}
		}
	}(stopCh, &wg)

	time.Sleep(1 * time.Second)
	close(stopCh)
	wg.Wait()
}
