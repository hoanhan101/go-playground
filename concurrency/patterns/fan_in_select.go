//
// Select
// Provide another way to handle multiple channels.
// It's like a swtich, but each case is a communication.
// - All channels are evaluated
// - Connection blocks until one communication can proceed, which then does.
// - If multiple can proceed, select chooses pseudo-randomly.
// - A default clause, if present, executes immediately if no channel is ready.
//

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Hoanh"), boring("An"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're boring. I am leaving!")
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

// FanIn using select
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}
