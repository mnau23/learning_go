package main

import "fmt"

func main() {
	ids := []int{23, 89, 65, 7, 11, 40}

	// Loop through
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i+1, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Addition
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	// Range with map
	emails := map[string](string){"A": "a@example.com", "B": "b@example.com", "C": "c@example.com"}

	for k, v := range emails {
		fmt.Printf("Name: %s; Email: %s\n", k, v)
	}
}
