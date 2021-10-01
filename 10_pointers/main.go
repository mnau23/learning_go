package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Printf("a: %d, b: %p, *b: %d\n", a, b, *b)
	fmt.Printf("%T\n", b)

	// Change value
	// (Everything in Go is pass-by-value)
	*b = 10
	fmt.Printf("a: %d\n", a)
}
