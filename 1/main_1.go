package main

import (
	"1/services"
	"fmt"
)

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	fmt.Println("numDecimal:", services.GetType(numDecimal))
	fmt.Println("numOctal:", services.GetType(numOctal))
	fmt.Println("numHexadecimal:", services.GetType(numHexadecimal))
	fmt.Println("pi:", services.GetType(pi))
	fmt.Println("name:", services.GetType(name))
	fmt.Println("isActive:", services.GetType(isActive))
	fmt.Println("complexNum:", services.GetType(complexNum))

	concatenated := services.Concatenate(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println("concatenated string:", concatenated)

	runeSlice := []rune(concatenated)
	fmt.Println("rune slice:", string(runeSlice))

	salt := "go-2024"
	hashed := services.Hash(concatenated, salt)
	fmt.Println("hashed:", hashed)
}
