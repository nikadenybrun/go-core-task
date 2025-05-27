package main

import (
	"3/services"
	"fmt"
)

func main() {
	originalSlice := services.GenerateSlice(10)
	fmt.Println("Original Slice:", originalSlice)

	evenSlice := services.SliceDer2(originalSlice)
	fmt.Println("Changed slice:", evenSlice)

	newSlice := services.AddElements(originalSlice, 42)
	fmt.Println("Slice after adding:", newSlice)

	copiedSlice := services.CopySlice(originalSlice)
	fmt.Println("Copied Slice:", copiedSlice)

	removedSlice := services.RemoveElement(originalSlice, 2)
	fmt.Println("Slice after removing element at index:", removedSlice)
}
