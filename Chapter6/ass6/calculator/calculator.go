package calculator

import "fmt"

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Cannot divide by zero")
		return 0
	}
	return a / b
}
