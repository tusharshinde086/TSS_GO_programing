package main

import "fmt"

func main() {
	var num1, num2, choice int

	fmt.Println("Enter two numbers for arithmatic operations  :")

	fmt.Scan(&num1, &num2)

	fmt.Println("Choose one operation to perform : \n 1. Add  2. Subtract 3. Multiply 4. Divide")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("Result: %d\n", num1+num2)
	case 2:
		fmt.Printf("Result: %d\n", num1-num2)
	case 3:
		fmt.Printf("Result: %d\n", num1*num2)
	case 4:
		fmt.Printf("Result: %.2f\n", float64(num1)/float64(num2))

	}
}
