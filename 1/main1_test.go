package main1_test

import (
	"1/services"
	"testing"
)

func TestGetType(t *testing.T) {
	var num int = 42
	expected := "int"
	if got := services.GetType(num); got != expected {
		t.Errorf("GetType(%v) = %v; want %v", num, got, expected)
	}
}

func TestConcatenate(t *testing.T) {
	result := services.Concatenate(1, 2.5, "test")
	expected := "12.5test"
	if result != expected {
		t.Errorf("Concatenate() = %v; want %v", result, expected)
	}
}

func TestHash(t *testing.T) {
	input := "Golang"
	salt := "go-2024"
	expected := "ac0e1077585ac7aaf393c53d7776282e7fdff694cb95ccfd2b517b76fbf49db3"
	result := services.Hash(input, salt)
	if result != expected {
		t.Errorf("Hash() = %v; want %v", result, expected)
	}
}
