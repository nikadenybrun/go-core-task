package services

import (
	"sync"
)

func MergeChannels(channels []chan int) chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for val := range c {
				merged <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
