package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int)   // Allocate memory for an integer using new
	one(xPtr)          // Pass the pointer to the function
	fmt.Println(*xPtr) // Dereference to print the value, which is now 1
}
