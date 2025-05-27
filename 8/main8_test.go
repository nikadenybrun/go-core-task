package main8_test

import (
	"8/services"
	"testing"
	"time"
)

func TestCustomWaitGroup1(t *testing.T) {
	wg := services.NewWaitGroup()
	const goroutines = 5
	counter := 0
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++
		}()
	}

	wg.Wait()

	if counter != goroutines {
		t.Errorf("expected %d, got %d", goroutines, counter)
	}
}

func TestCustomWaitGroup2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, but did not")
		}
	}()

	wg := services.NewWaitGroup()
	wg.Add(1)
	wg.Done()
	wg.Done()
}

func TestCustomWaitGroup3(t *testing.T) {
	wg := services.NewWaitGroup()

	const goroutines = 3
	wg.Add(goroutines)

	counter := 0

	for i := 0; i < goroutines; i++ {
		go func() {
			time.Sleep(10 * time.Millisecond)
			defer wg.Done()
			counter++
		}()
	}

	before := counter
	wg.Wait()
	after := counter

	if before != 0 {
		t.Errorf("counter must be 0, got %d", before)
	}
	if after != goroutines {
		t.Errorf("expected %d, got %d", goroutines, after)
	}
}
