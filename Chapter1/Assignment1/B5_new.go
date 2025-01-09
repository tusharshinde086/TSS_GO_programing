package main

import "fmt"

func main() {

	ptr := new(int)

	*ptr = 100

	fmt.Println("Value:", *ptr)
	fmt.Println("Address:", ptr)
}
