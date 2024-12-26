package main

import "fmt"

func main() {
	var ip *int
	var fp *float32

	var a = 10
	var b = 3.14

	ip = &a
	fp = &b

	fmt.Println("Value of a:", *ip)
	fmt.Println("Value of b:", *fp)

	fmt.Println("Address of a:", ip)
	fmt.Println("Address of b:", fp)
}
