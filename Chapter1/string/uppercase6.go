package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Println("Enter a string:")
	fmt.Scan(&str)

	fmt.Println("Uppercase string:", strings.ToUpper(str))
}
