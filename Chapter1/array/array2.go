package main

import (
	"fmt"
)

func main() {
	var rows, cols int

	// Get dimensions from the user
	fmt.Print("Enter number of rows and columns: ")
	fmt.Scan(&rows, &cols)

	// Declare a 2D array based on input size
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	// Take input for each element
	fmt.Println("Enter the elements row by row:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	// Display the 2D array
	fmt.Println("Your matrix is:")
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}

// //  go run array2.go
// Enter number of rows and columns: 2 3
// Enter the elements row by row:
// 12 13 14
// 15 16 17
// Your matrix is:
// 12 13 14
// 15 16 17
