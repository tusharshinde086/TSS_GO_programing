package main

import "fmt"

func main() {
	num := 10

	if num > 5 {
		fmt.Println("Number is greater than 5")
	}

	if num%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 7; num < 10 { // Short statement in if
		fmt.Println("Number is less than 10")
	}
}

// Number is greater than 5
// Number is even
// Number is less than 10
