package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Enter a string with spaces:")
	fmt.Scan(&str)

	result := strings.TrimSpace(str)
	fmt.Println("Trimmed string:", result)
}
