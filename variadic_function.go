package main

import "fmt"

func main() {
    total := sum(10, 10, 10)
    fmt.Println(total)
}

func sum(numbers ...int) int {
    total := 0
    for _, v := range numbers {
        total += v
    }
    return total
}
