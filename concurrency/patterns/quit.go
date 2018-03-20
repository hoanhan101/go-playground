//
// Quit channel (not working yet)
//

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("Hoanh")
	quit := make(chan string)

	for i := rand.Intn(10); i > 0; i-- {
		fmt.Println(<-c)
	}

	quit <- "Bye"

	fmt.Printf("%q\n", <-quit)
}

func boring(msg string) <-chan string {
	c := make(chan string)

	go func() {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-quit:
			cleanup()
			quit <- "See you!"
			return
		}
	}()

	return c
}
