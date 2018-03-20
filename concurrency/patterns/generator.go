//
// Generator: function that returns a channel
//

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	hoanh := boring("Hoanh")
	an := boring("An")

	for i := 0; i < 5; i++ {
		fmt.Println(<-hoanh)
		fmt.Println(<-an)
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
