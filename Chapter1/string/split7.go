package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Enter a string:")
	fmt.Scan(&str)

	words := strings.Fields(str)
	fmt.Println("Words in the string:", words)
}
