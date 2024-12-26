package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	isPrime := true
	if num <= 1 {
		isPrime = false
	} else {
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Println(num, "is a Prime number")
	} else {
		fmt.Println(num, "is not a Prime number")
	}
}
