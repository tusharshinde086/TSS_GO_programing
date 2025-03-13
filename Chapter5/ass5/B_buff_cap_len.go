package main

import (
	"fmt"
)

func main() {

	myChnl := make(chan int, 5)

	myChnl <- 10
	myChnl <- 20
	myChnl <- 30

	fmt.Println("Capacity of the channel:", cap(myChnl)) // Should print 5
	fmt.Println("Length of the channel:", len(myChnl))   // Should print 3

	for i := 0; i < 3; i++ {
		value := <-myChnl
		fmt.Println("Read from channel:", value)
	}

	fmt.Println("Length of the channel after reading:", len(myChnl)) // Should print 0
}
