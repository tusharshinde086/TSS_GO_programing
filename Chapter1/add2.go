package main

import "fmt"

func main() {
    var a, b, c float64
    fmt.Println("Enter two numbers for addition:")
    fmt.Scanln(&a, &b)
    c = a + b
    fmt.Println("The sum is:", c)
}

