package main8_test

import (
	"9/services"
	"sync"
	"testing"
)

func TestCubePipeline(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(1)
	go services.Cube(in, out, &wg)

	go func() {
		defer close(in)
		inputs := []uint8{1, 2, 3, 4, 5}
		for _, v := range inputs {
			in <- v
		}
	}()

	expected := []float64{1, 8, 27, 64, 125}
	results := []float64{}
	for v := range out {
		results = append(results, v)
	}

	wg.Wait()

	if len(results) != len(expected) {
		t.Errorf("expected length %d, got %d", len(expected), len(results))
	}
	for i := range expected {
		if results[i] != expected[i] {
			t.Errorf("index %d: expected %f, got %f", i, expected[i], results[i])
		}
	}
}

func TestCube(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(1)
	go services.Cube(in, out, &wg)
	close(in)
	count := 0
	for range out {
		count++
	}

	wg.Wait()

	if count != 0 {
		t.Errorf("expected 0, got %d", count)
	}
}
