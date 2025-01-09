package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter a string:")
	fmt.Scan(&str)

	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}

	if str == reversed {
		fmt.Println("The string is a palindrome.")
	} else {
		fmt.Println("The string is not a palindrome.")
	}
}
