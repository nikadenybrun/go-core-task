package main4_test

import (
	"5/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInnerSlices1(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	expected := []int{3, 64}
	boolean, res := services.InnerSlices(a, b)

	assert.Equal(t, expected, res)
	assert.True(t, boolean)
}
func TestInnerSlices2(t *testing.T) {
	slice1 := []int{}
	slice2 := []int{64, 3, 2}
	expected := []int{}
	boolean, res := services.InnerSlices(slice1, slice2)

	assert.Equal(t, expected, res)
	assert.False(t, boolean)
}
func TestInnerSlices3(t *testing.T) {
	slice1 := []int{64, 3, 2}
	slice2 := []int{}
	expected := []int{}
	boolean, res := services.InnerSlices(slice1, slice2)

	assert.Equal(t, expected, res)
	assert.False(t, boolean)
}
func TestInnerSlices4(t *testing.T) {
	slice1 := []int{64, 3, 2}
	slice2 := []int{64, 3, 2}
	expected := []int{64, 3, 2}
	boolean, res := services.InnerSlices(slice1, slice2)

	assert.Equal(t, expected, res)
	assert.True(t, boolean)
}
func TestInnerSlices5(t *testing.T) {
	slice1 := []int{}
	slice2 := []int{}
	expected := []int{}
	boolean, res := services.InnerSlices(slice1, slice2)

	assert.Equal(t, expected, res)
	assert.False(t, boolean)
}
func TestInnerSlices6(t *testing.T) {
	slice1 := []int{64, 3, 2}
	slice2 := []int{64, 3, 2, 64, 3, 2}
	expected := []int{64, 3, 2}
	boolean, res := services.InnerSlices(slice1, slice2)

	assert.Equal(t, expected, res)
	assert.True(t, boolean)
}
func TestInnerSlices7(t *testing.T) {
	slice1 := []int{64, 3, 2, 64, 3, 2}
	slice2 := []int{64, 3, 2, 64, 3, 2}
	expected := []int{64, 3, 2, 64, 3, 2}
	boolean, res := services.InnerSlices(slice1, slice2)

	assert.Equal(t, expected, res)
	assert.True(t, boolean)
}
func TestInnerSlices8(t *testing.T) {
	slice1 := []int{64, 3, 2, 64, 3, 2}
	slice2 := []int{64, 3, 64, 3, 6}
	expected := []int{64, 3, 64, 3}
	boolean, res := services.InnerSlices(slice1, slice2)

	assert.Equal(t, expected, res)
	assert.True(t, boolean)
}
