package main

import "fmt"

// Arrays: must be fixed length and have a type
// Slice: an array that doesn't have a fixed type

func main() {
	// === Arrays ===
	var fruits [3]string

	// Assign values
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Grape"

	fmt.Println(fruits)
	fmt.Println(fruits[1])
	fmt.Println(fruits[1:2]) // like [i:j)

	// Declare and assign values
	dessert := [2]string{"Cookie", "Cake"}

	fmt.Println(len(dessert))
	fmt.Println(dessert[1])

	// === Slices ===
	meal := []string{"Pizza", "Pasta"}
	fmt.Println(len(meal))
}
