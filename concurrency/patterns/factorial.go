package main

import (
	"fmt"
	"sync"
)

func main() {
	in := generate(10)

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is
	// closed.
	// Distribute work arcoss multiple functions, 10 goroutines, that all read
	// from in
	c0 := factorial(in)
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)
	c6 := factorial(in)
	c7 := factorial(in)
	c8 := factorial(in)
	c9 := factorial(in)

	// FAN IN
	// Multiplex multiple channels onto a single channel
	// Merge the channels from c0 to c9 onto a single channel
	for n := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		fmt.Println("main()", n)
	}
}

func generate(n int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			fmt.Println("generate()", i)
			c <- i
		}
		close(c)
	}()
	return c
}

func factorial(in <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for n := range in {
			fmt.Println("factorial()", n, fact(n))
			c <- fact(n)
		}
		close(c)
	}()
	return c
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			// fmt.Println("output()", n)
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are done.
	// This must start after the wg.Add() call
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
