package main

import (
	"fmt"
)

func main() {
	// 1. Creating a slice
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println("Original Slice:", numbers)

	// 2. Append elements to a slice
	numbers = append(numbers, 60, 70)
	fmt.Println("After Append:", numbers)

	// 3. Removing an element (e.g., remove element at index 2)
	index := 2
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Println("After Removing Index 2:", numbers)

	// 4. Copying a slice
	newSlice := make([]int, len(numbers))
	copy(newSlice, numbers)
	fmt.Println("Copied Slice:", newSlice)

	// 5. Slicing a slice (getting a sub-slice)
	subSlice := numbers[1:4] // Includes index 1 to 3
	fmt.Println("Sub-slice (1 to 3):", subSlice)
}
