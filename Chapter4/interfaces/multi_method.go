package main

import "fmt"

// Define an interface with multiple methods
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	var s Shape = Rectangle{width: 10, height: 5}
	fmt.Println("Area:", s.Area())           // Output: Area: 50
	fmt.Println("Perimeter:", s.Perimeter()) // Output: Perimeter: 30
}
