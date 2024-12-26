//2. WAP to check whether a number is even or odd

package main

import "fmt"

func main() {
	// Declare a number
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	// Check if the number is even or odd
	if num%2 == 0 {
		fmt.Println("The number", num, "is Even.")
	} else {
		fmt.Println("The number", num, "is Odd.")
	}
}
