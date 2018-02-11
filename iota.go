package main

import "fmt"

const (
    // Start with 0, automatically increase by 1 everytime it get called
    _ = iota 

    // Shift 10 place to the left to get KB
    // Shift 20 place to the left to get MB
    // ...
    KB = 1 << (iota * 10)
    MB = 1 << (iota * 10)
    GB = 1 << (iota * 10)
    TB = 1 << (iota * 10)
)

func main() {
    fmt.Printf("Binary: %b, Decimal: %d\n", KB, KB)
    fmt.Printf("Binary: %b, Decimal: %d\n", MB, MB)
    fmt.Printf("Binary: %b, Decimal: %d\n", GB, GB)
    fmt.Printf("Binary: %b, Decimal: %d\n", TB, TB)
}
