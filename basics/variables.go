package main

import "fmt"

func main() {
	// Shorthand for declaring and initializing variables
	a := 1
	b := "string"
	c := true
	d := 1.0

	fmt.Printf("Default value is %v, which has type %T\n", a, a)
	fmt.Printf("Default value is %v, which has type %T\n", b, b)
	fmt.Printf("Default value is %v, which has type %T\n", c, c)
	fmt.Printf("Default value is %v, which has type %T\n", d, d)
}
