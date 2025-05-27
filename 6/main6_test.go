package main6_test

import (
	"6/services"
	"testing"
)

func TestRandomGenerator(t *testing.T) {
	count := 10
	ch := make(chan int)

	go services.StartRandomGenerator(count, ch)

	nums := make([]int, 0, count)
	for num := range ch {
		nums = append(nums, num)
	}

	if len(nums) != count {
		t.Errorf("expected %d numbers but got %d", count, len(nums))
	}
}
