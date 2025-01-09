package main

import "fmt"

func main() {
	var num int = 65
	var ch byte = byte(num) // Casting from int to byte (which is a character)
	fmt.Println("Int to byte (char):", ch)
}
