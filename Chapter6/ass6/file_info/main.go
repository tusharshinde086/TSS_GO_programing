package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("data.xml")
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fmt.Printf("File Name: %s\n", fileInfo.Name())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Is Directory: %t\n", fileInfo.IsDir())
	fmt.Printf("Last Modified: %s\n", fileInfo.ModTime())
}
