package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Enter a sentence to change in words:")
	fmt.Scan(&str)

	words := strings.Fields(str)
	fmt.Println("The words in the string are:", words)
}
