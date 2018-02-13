package main

import "fmt"

func main() {
    helloWorld := func() {
        fmt.Println("Hello World")
    }

    helloWorld()

    // Wrapper
    increment := wrapper()
    fmt.Println(increment())
    fmt.Println(increment())
}

func wrapper() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}
