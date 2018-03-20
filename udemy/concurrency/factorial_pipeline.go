package main

import (
    "fmt"
    "math/rand"
)

func main() {
    in := generate(10, 20)
    f := factorial(in)
    for n := range f {
        fmt.Println(n)
    }
}

func generate(n int, r int) <-chan int {
    out := make(chan int)
    go func() {
        for i := 0; i < n; i++ {
            out <- rand.Intn(r)
        }
        close(out)
    }()
    return out
}

func factorial(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for i := range in {
            out <- fact(i)
        }
        close(out)
    }()
    return out
}

func fact(n int) int {
    total := 1
    for i := n; i > 0; i-- {
        total *= i
    }
    return total
}
