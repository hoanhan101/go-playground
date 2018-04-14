package main

import "fmt"

func main() {
	// Variadic parameters
	total := sum(10, 10, 10)
	fmt.Println(total)

	// Variadic arguments
	numbers := []int{10, 10, 10}
	result := sum(numbers...)
	fmt.Println(result)
}

func sum(numbers ...int) int {
	total := 0
	for _, v := range numbers {
		total += v
	}
	return total
}
