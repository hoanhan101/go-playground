package main

import (
    "fmt"
)

func main() {
    c := fanIn(hello("Hoanh", 5), hello("An", 10))
    for i := 0; i < 20; i++ {
        fmt.Println(<-c)
    }
}

func hello(name string, n int) <-chan string {
    c := make(chan string)
    go func() {
        for i := 0; i < n; i++ {
            c <- fmt.Sprintf("%v %v", name, i)
        }
        close(c)
    }()
    return c
}

func fanIn(c1, c2 <- chan string) <-chan string {
    c := make(chan string)
    go func() {
        for {
            c <- <-c1
        }
    }()
    go func() {
        for {
            c <- <-c2
        }
    }()
    return c
}
