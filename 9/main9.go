package main

import (
	"9/services"
	"fmt"
	"sync"
)

func main() {
	in := make(chan uint8)
	out := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(1)
	go services.Cube(in, out, &wg)

	go func() {
		for i := uint8(1); i <= 5; i++ {
			in <- i
		}
		close(in)
	}()

	for result := range out {
		fmt.Printf("%.2f\n", result)
	}

	wg.Wait()
}
