package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchData(source string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal completion
	fmt.Printf("Fetching data from %s...\n", source)
	time.Sleep(2 * time.Second) // Simulate network delay
	fmt.Printf("Data fetched from %s\n", source)
}

func main() {
	var wg sync.WaitGroup // Create WaitGroup
	sources := []string{"Source 1", "Source 2", "Source 3"}

	for _, source := range sources {
		wg.Add(1)                 // Increment counter
		go fetchData(source, &wg) // Start goroutine
	}

	wg.Wait() // Wait for all to finish
	fmt.Println("All data fetched.")
}
