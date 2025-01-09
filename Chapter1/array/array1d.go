package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number of elements:")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Enter the elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("The array is:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
}
