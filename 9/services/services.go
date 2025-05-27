package services

import (
	"sync"
)

func Cube(in <-chan uint8, out chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		fl := float64(num)
		out <- fl * fl * fl
	}
	close(out)
}
