package main

import "fmt"

// Add function to add two integers
func Add(a int, b int) int {
	return a + b
}

func main() {
	result := Add(3, 4)
	fmt.Printf("3 + 4 = %d\n", result)
}
