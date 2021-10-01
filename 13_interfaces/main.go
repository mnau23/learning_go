package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

// Define structs
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// Methods
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{1, 1, 5}
	rectangle := Rectangle{3, 3}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
