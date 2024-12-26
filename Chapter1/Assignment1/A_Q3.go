//3. WAP to swap two numbers without using a temporary variable

package main

import "fmt"

func main() {
	// Declare two numbers
	var a, b int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	// Swap the numbers without using a temporary variable
	a = a + b
	b = a - b
	a = a - b

	// Print swapped values
	fmt.Println("After swapping:")
	fmt.Println("First number:", a)
	fmt.Println("Second number:", b)
}
