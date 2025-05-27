package main4_test

import (
	"4/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeSlices1(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}
	res := services.ChangeSlices(slice1, slice2)

	assert.Equal(t, expected, res)
}
func TestChangeSlices2(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{}
	res := services.ChangeSlices(slice1, slice2)

	assert.Equal(t, expected, res)
}
func TestChangeSlices3(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{}
	expected := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	res := services.ChangeSlices(slice1, slice2)

	assert.Equal(t, expected, res)
}
func TestChangeSlices4(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	expected := []string{}
	res := services.ChangeSlices(slice1, slice2)

	assert.Equal(t, expected, res)
}
func TestChangeSlices5(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{}
	expected := []string{}
	res := services.ChangeSlices(slice1, slice2)

	assert.Equal(t, expected, res)
}
func TestChangeSlices6(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry"}
	slice2 := []string{"apple", "banana", "cherry", "apple", "banana", "cherry"}
	expected := []string{}
	res := services.ChangeSlices(slice1, slice2)

	assert.Equal(t, expected, res)
}
func TestChangeSlices7(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "apple", "banana", "cherry"}
	slice2 := []string{"apple", "banana", "cherry"}
	expected := []string{}
	res := services.ChangeSlices(slice1, slice2)

	assert.Equal(t, expected, res)
}
func TestChangeSlices8(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "apple", "banana", "cherry"}
	slice2 := []string{"apple", "banana"}
	expected := []string{"cherry", "cherry"}
	res := services.ChangeSlices(slice1, slice2)

	assert.Equal(t, expected, res)
}
