package main

import "fmt"

func main() {
    slice := make([]int, 0, 5)
    fmt.Println(slice)
    fmt.Println(len(slice))
    fmt.Println(cap(slice))

    // If the length of slice exceeds its capacity,
    // it automatically double the capacity everytime
    for i := 0; i < 100; i++ {
        slice = append(slice, i)
        fmt.Println(i, len(slice), cap(slice))
    }


    // Shorthand
    mySlice := []int{}

    // Cannot do this mySlice[0] = 1. Must do this:
    mySlice = append(mySlice, 1)
    fmt.Println(mySlice)

    // However, make function reserves spot so you can do that
    // In our example, make a slice of 5 slots out of 10-slot underlying array
    // This is preferable
    // Slice is a reference type, points to the underlying array
    anotherSlice := make([]int, 5, 10)
    anotherSlice[0] = 1
    fmt.Println(anotherSlice)
}
