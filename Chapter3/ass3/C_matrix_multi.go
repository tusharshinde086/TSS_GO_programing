package main

import (
	"fmt"
)

// Function to accept a matrix from user
func inputMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			fmt.Printf("Enter element [%d][%d]: ", i+1, j+1)
			fmt.Scan(&matrix[i][j])
		}
	}
	return matrix
}

// Function to display a matrix
func displayMatrix(matrix [][]int) {
	for _, row := range matrix {
		for _, val := range row {
			fmt.Printf("%d\t", val)
		}
		fmt.Println()
	}
}

// Function to multiply two matrices
func multiplyMatrices(A, B [][]int, r1, c1, c2 int) [][]int {
	result := make([][]int, r1)
	for i := range result {
		result[i] = make([]int, c2)
		for j := range result[i] {
			for k := 0; k < c1; k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return result
}

func main() {
	var r1, c1, r2, c2 int

	// Accept dimensions of the matrices
	fmt.Print("Enter rows and columns of first matrix: ")
	fmt.Scan(&r1, &c1)

	fmt.Print("Enter rows and columns of second matrix: ")
	fmt.Scan(&r2, &c2)

	// Matrix multiplication condition: c1 must be equal to r2
	if c1 != r2 {
		fmt.Println("Matrix multiplication not possible! Columns of first matrix must equal rows of second matrix.")
		return
	}

	// Accept matrices
	fmt.Println("Enter elements for first matrix:")
	A := inputMatrix(r1, c1)

	fmt.Println("Enter elements for second matrix:")
	B := inputMatrix(r2, c2)

	// Perform multiplication
	result := multiplyMatrices(A, B, r1, c1, c2)

	// Display results
	fmt.Println("\nFirst Matrix:")
	displayMatrix(A)

	fmt.Println("\nSecond Matrix:")
	displayMatrix(B)

	fmt.Println("\nResultant Matrix (Multiplication):")
	displayMatrix(result)
}
