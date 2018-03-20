package main

import (
	"fmt"
)

func main() {
    c := incrementor()
    cSum := puller(c)

    for i := range cSum {
        fmt.Println(i)
    }
}

// Receive-only channel
func incrementor() <-chan int {
    out := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            out <- i
        }
        close(out)
    }()
    return out
}

func puller(c <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        var sum int
        for i := range c {
            sum += i
        }
        out <- sum
        close(out)
    }()
    return out
}
