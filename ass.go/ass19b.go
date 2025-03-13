package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file in read-only mode
	file, err := os.Open("example.txt") // Replace "example.txt" with your file name
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed after we're done

	// Read the content of the file (for demonstration purposes)
	content := make([]byte, 100) // Buffer to hold file content
	n, err := file.Read(content)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the content read from the file
	fmt.Printf("Read %d bytes: %s\n", n, string(content[:n]))
}
