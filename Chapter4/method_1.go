package main

import "fmt"

func main() {
	print(area, 2, 4)
	print(sum, 2, 4)
}
func print(f func(int, int) int, a, b int) {
	fmt.Println(f(a, b))
}
func area(a, b int) int {
	return a * b
}
func sum(a, b int) int {
	return a + b
}
