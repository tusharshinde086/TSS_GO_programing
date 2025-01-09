package main

import "fmt"

func main() {
	var pi float32 = 3.14
	var num int = int(pi) // Casting from float32 to int
	fmt.Println("Float to int:", num)
}
