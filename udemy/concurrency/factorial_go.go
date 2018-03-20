package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    for i := 0; i < 1000; i++ {
        go func() {
            f := factorial(rand.Intn(100))
            for n := range f {
                fmt.Println(n)
            }
        }()
    }
    time.Sleep(3 * time.Millisecond)
}

func factorial(n int) chan int {
    out := make(chan int)

    go func() {
        total := 1
        for i := n; i > 0; i-- {
            total *= i
        }
        out <- total
        close(out)
    }()
    return out
}
