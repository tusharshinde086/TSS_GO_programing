package main

import "fmt"

func main() {
	var a float64 = 9.99
	var b int = int(a) // Explicit casting: float64 to int
	fmt.Println("Explicit casting:", b)
}
