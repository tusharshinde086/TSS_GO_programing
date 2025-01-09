package main

import "fmt"

func main() {
	var str string
	fmt.Println("Enter a string to count vowels:")
	fmt.Scan(&str)

	count := 0
	for _, char := range str {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
			char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			count++
		}
	}

	fmt.Println("Number of vowels:", count)
}
