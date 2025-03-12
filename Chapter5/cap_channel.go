package main

import (
	"fmt"
)

func main() {
	// Create a buffered channel with capacity 5
	myChnlCap := make(chan string, 5)

	// Send values to the channel
	myChnlCap <- "X"
	myChnlCap <- "Y"
	myChnlCap <- "Z"

	// Print the capacity of the channel
	fmt.Println("Capacity of the channel (cap):", cap(myChnlCap)) // Should print 5
}
