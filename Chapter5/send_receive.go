package main

import (
	"fmt"
	"time"
)

func unbufferedChannel() {
	ch := make(chan int) // Create an unbuffered channel

	go func() {
		fmt.Println("Sending value 42")
		ch <- 42 // Send value to the channel
		fmt.Println("Value sent")
	}()

	fmt.Println("Receiving value")
	value := <-ch // Receive value from the channel
	fmt.Println("Received:", value)
}

func bufferedChannel() {
	ch := make(chan int, 3) // Create a buffered channel with a capacity of 3

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("Sending value:", i)
			ch <- i // Send value to the channel
			fmt.Println("Value sent:", i)
		}
		close(ch) // Close the channel when done
	}()

	for value := range ch {
		fmt.Println("Received:", value) // Print received values
	}
}

func main() {
	fmt.Println("Unbuffered Channel Example:")
	unbufferedChannel()

	time.Sleep(time.Second) // Wait for a moment before the next example

	fmt.Println("\nBuffered Channel Example:")
	bufferedChannel()
}
