package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()

	// Adds the next number in the loop
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
