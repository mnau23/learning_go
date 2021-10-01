package main

import "fmt"

func main() {
	// Define map
	// key: first name, value: email
	// make(map[key_type]value_type)
	emails := make(map[string]string)

	// Assign k-v
	emails["A"] = "a@example.com"
	emails["B"] = "b@example.com"
	emails["C"] = "c@example.com"

	// Define and assign
	// emails := map[string](string){"A": "a@example.com", "B": "b@example.com", "C": "c@example.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["B"])

	// Delete
	delete(emails, "A")
	fmt.Println(emails)
}
