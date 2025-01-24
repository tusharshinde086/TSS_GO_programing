package main

import "fmt"

func main() {
	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println("slice A:", a)
	fmt.Printf("capacity is %d and length is %d\n", cap(a), len(a))
	a = append(a, 30, 40, 50, 60, 70)
	fmt.Println("slice A after append:", a)
	fmt.Printf("capacity after add %d and length is %d\n", cap(a), len(a))
}
