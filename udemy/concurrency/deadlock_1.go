package main

import "fmt"

func main() {
	c := make(chan int)

	// Send 1 to the channel.
	// If we do c <- 1 alone, we will create a deadlock.
	// The solution is to wrap it within a go func().
	go func() {
		c <- 1
	}()

	// Read from the channel
	fmt.Println(<-c)
}
