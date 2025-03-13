package main

import "fmt"

// Add function to add two integers
func Add(a int, b int) int {
	return a + b
}

func main() {
	a, b := 5, 10
	result := Add(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, result)
}
