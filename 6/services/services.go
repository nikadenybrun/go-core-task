package services

import (
	"math/rand"
	"sync"
	"time"
)

func RandomGenerator(count int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		ch <- rand.Intn(100)
	}
}

func StartRandomGenerator(count int, ch chan<- int) {
	var wg sync.WaitGroup
	wg.Add(1)

	go RandomGenerator(count, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()
}
