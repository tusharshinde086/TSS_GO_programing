package main

import "fmt"

// Function that tries to modify a value (call by value)
func modifyValue(a int) {
	a = a + 10
	fmt.Println("Inside modifyValue function:", a)
}

func main() {
	num := 5
	fmt.Println("Before modifyValue:", num)
	modifyValue(num)
	fmt.Println("After modifyValue:", num)
}
