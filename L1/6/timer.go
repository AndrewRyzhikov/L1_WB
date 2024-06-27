package five

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	timer := time.NewTimer(2 * time.Second)
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-timer.C:
				return
			default:
				fmt.Println("Working...")
				time.Sleep(1 * time.Second)
			}
		}
	}(&wg)

	time.Sleep(3 * time.Second)
	wg.Wait()
}
