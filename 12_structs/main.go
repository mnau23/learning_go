package main

import (
	"fmt"
	"strconv"
)

// Define struct
type Person struct {
	name, surname string
	age           int
	country       string
}

// Methods go outside of the struct

// Value receiver
func (p Person) greet() string {
	return "Hello, my name is " + p.name + " " + p.surname + " and I am " + strconv.Itoa(p.age)
}

// Pointer receiver
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// Init person
	person1 := Person{name: "N1", surname: "S1", age: 99, country: "Italy"}
	person2 := Person{"N2", "S2", 98, "Italy"}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person2.age)
	person1.hasBirthday()
	fmt.Println(person1.greet())
}
