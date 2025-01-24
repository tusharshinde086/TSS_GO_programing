package main

import "fmt"

type Book struct {
	BookID string
	Title  string
	Author string
	Price  float64
}

func main() {
	var n int
	fmt.Print("Enter the number of books: ")
	fmt.Scan(&n)

	books := make([]Book, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter details for Book %d\n", i+1)
		fmt.Print("BookID: ")
		fmt.Scan(&books[i].BookID)
		fmt.Print("Title: ")
		fmt.Scan(&books[i].Title)
		fmt.Print("Author: ")
		fmt.Scan(&books[i].Author)
		fmt.Print("Price: ")
		fmt.Scan(&books[i].Price)
	}

	fmt.Println("\nBook Details:")
	for i := 0; i < n; i++ {
		fmt.Printf("Book %d:\n", i+1)
		fmt.Printf("BookID: %s\n", books[i].BookID)
		fmt.Printf("Title: %s\n", books[i].Title)
		fmt.Printf("Author: %s\n", books[i].Author)
		fmt.Printf("Price: %.2f\n\n", books[i].Price)
	}
}
