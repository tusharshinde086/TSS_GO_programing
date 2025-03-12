package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int) // Create an unbuffered channel

	go func() {
		fmt.Println("Sending value 42")
		ch <- 42 // Send value to the channel (blocks until received)
		fmt.Println("Value sent")
	}()

	time.Sleep(time.Second) // Simulate some work in the main goroutine

	fmt.Println("Receiving value")
	value := <-ch // Receive value from the channel (blocks until sent)
	fmt.Println("Received:", value)
}
