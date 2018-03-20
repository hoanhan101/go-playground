package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// If we only read from <- c, then it only return 0.
	// In order to read everything from the channel, we use range.
	for i := range c {
		fmt.Println(i)
	}
}
