package services

import (
	"time"

	"golang.org/x/exp/rand"
)

func GenerateSlice(size int) []int {
	rand.Seed(uint64(time.Now().UnixNano()))
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100)
	}
	return slice
}

func SliceDer2(slice []int) []int {
	var newSlice []int
	for _, v := range slice {
		if v%2 == 0 {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func AddElements(slice []int, element int) []int {
	return append(slice, element)
}

func CopySlice(slice []int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func RemoveElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
