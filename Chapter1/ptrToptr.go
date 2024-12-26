package main

import "fmt"

func main() {
	var a = 10
	var ptr *int = &a         // ptr is a pointer to an integer
	var ptrToPtr **int = &ptr // ptrToPtr is a pointer to a pointer to an integer

	fmt.Println("Value of a:", a)
	fmt.Println("Value pointed by ptr:", *ptr)
	fmt.Println("Value pointed by ptrToPtr:", **ptrToPtr) // Dereferencing twice to get the value of a
}
