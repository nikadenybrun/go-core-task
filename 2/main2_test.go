package main3_test

import (
	"3/services"
	"reflect"
	"testing"
)

func TestGenerateSlice(t *testing.T) {
	slice := services.GenerateSlice(10)
	if len(slice) != 10 {
		t.Errorf("expected slice length of 10, got %d", len(slice))
	}
}

func TestSliceDer2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{2, 4, 6}
	result := services.SliceDer2(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3}
	element := 4
	expected := []int{1, 2, 3, 4}
	result := services.AddElements(input, element)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3}
	copied := services.CopySlice(input)
	copied[0] = 99 // Изменяем копию
	if input[0] == 99 {
		t.Error("Original slice should not be affected by changes in the copied slice")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{1, 2, 4}
	result := services.RemoveElement(input, 2) // Удаляем элемент с индексом 2
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
