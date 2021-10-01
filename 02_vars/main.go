package main

import "fmt"

var country string = "Italy"

// country := "Italy" >>> not possible outside a function

func main() {
	// Main types:
	// string
	// bool
	// int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64, uintptr
	// byte (alias for uint8)
	// rune (alias for int32)
	// float32, float64
	// complex64, complex128

	//var name = "Pippo"
	name := "Pippo"
	var age int = 10
	height := 1.5

	fmt.Println(name, age, country)
	fmt.Printf("%T\n", height)

	var flag = true
	fmt.Println(flag)
	flag = false
	fmt.Println(flag)

	const flag2 = false
	// flag2 = true >>> returns an error

	word1, word2 := "Example", "Something"
	fmt.Println(word1, word2)
}
