package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var count int
	fmt.Println("Enter a string to repeat:")
	fmt.Scan(&str)
	fmt.Println("Enter the number of repetitions:")
	fmt.Scan(&count)

	result := strings.Repeat(str, count)
	fmt.Println("Repeated string:", result)
}
