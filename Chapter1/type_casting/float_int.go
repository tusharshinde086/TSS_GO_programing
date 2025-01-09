package main

import "fmt"

func main() {
	var a float64 = 10.75
	var b int = int(a) // Casting from float64 to int
	fmt.Println("Casting float64 to int:", b)
}
