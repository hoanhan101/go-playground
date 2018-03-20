package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()

    f := factorial(20)

    for i := range f {
        t := time.Now()
        elapsed := t.Sub(start)
        fmt.Println(elapsed, i)
    }
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
