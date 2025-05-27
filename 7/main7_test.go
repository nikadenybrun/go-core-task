package main7_test

import (
	"7/services"
	"reflect"
	"sort"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 2
		close(ch1)
	}()

	go func() {
		ch2 <- 3
		ch2 <- 4
		close(ch2)
	}()

	go func() {
		ch3 <- 5
		ch3 <- 6
		close(ch3)
	}()

	merged := services.MergeChannels([]chan int{ch1, ch2, ch3})

	var result []int
	for val := range merged {
		result = append(result, val)
	}
	sort.Ints(result)
	expected := []int{1, 2, 3, 4, 5, 6}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v but got %v", expected, result)
	}
}
