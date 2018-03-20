package main

import (
    "fmt"
)

func main() {
    c := generate(2, 3)
    out := square(c)

    fmt.Println(<-out)
    fmt.Println(<-out)
}

func generate(nums ...int) chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(c chan int) chan int {
    out := make(chan int)
    go func() {
        for n := range c {
            out <- n * n
        }
        close(out)
    }()
    return out
}
