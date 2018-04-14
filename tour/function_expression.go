package main

import "fmt"

func main() {
	// Function expression
	helloWorld := func() {
		fmt.Println("Hello World")
	}

	helloWorld()

	// Wrapper
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())

	// Anonymous Self-Executing Function
	func() {
		fmt.Println("Hello World")
	}()
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
