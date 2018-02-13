package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    printOut(numbers, func(n int) {
        fmt.Println(n)
    })
}

// Callback: passing a function as an argument
func printOut(numbers []int, callback func(int)) {
    for _, n := range numbers {
        callback(n)
    }
}
