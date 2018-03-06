package main

import "fmt"

func main() {
    defer world()
    hello()
}

func world() {
    fmt.Println("World")
}

func hello() {
    fmt.Println("Hello")
}
