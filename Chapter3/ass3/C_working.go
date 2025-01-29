package main

import "fmt"

func main() {
	// Initial slice
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial slice:", slice)

	// Append elements to the slice
	slice = append(slice, 6, 7, 8)
	fmt.Println("Slice after append:", slice)

	// Remove an element at index 2 (remove 3)
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("Slice after removing element at index 2:", slice)

	// Copy slice
	copiedSlice := make([]int, len(slice))
	copy(copiedSlice, slice)
	fmt.Println("Copied slice:", copiedSlice)
}

