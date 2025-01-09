package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Enter a string with spaces:")
	fmt.Scan(&str)

	result := strings.ReplaceAll(str, " ", "")
	fmt.Println("String without spaces:", result)
}
