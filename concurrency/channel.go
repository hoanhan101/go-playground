package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cSemaphore()
}

func cUnbuffered() {
	// This is unbuffered channel
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// This stops right here until until something takes the value off,
			// which is what we are doing in the next goroutine function
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	// Wait for some time for these above functions to execute. Otherwise, need
	// to use WaitGroup
	time.Sleep(time.Second)

}

func cRange() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}

func cWaitGroup() {
	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}

func cSemaphore() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	// If we don't run this with goroutine, we will have a deadlock.
	// We will never get to the range c if we don't use goroutine.
	// If we don't get there, it is blocking so we hang there forever.
	go func() {
		<-done
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}

func cSemaphore2() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	for i := 0; i < n; i++ {
		go func() {
			for i := 0; i < n; i++ {
				c <- i
			}
			done <- true
		}()
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}
