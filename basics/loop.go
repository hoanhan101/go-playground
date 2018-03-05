package main

import "fmt"

func main() {
	// for loop
	for a := 0; a < 10; a++ {
		fmt.Println(a)
	}

	// while loop
	b := 0
	for b < 10 {
		fmt.Println(b)
		b++
	}

	// while True
	c := 0
	for {
		fmt.Println(c)
		if c >= 10 {
			break
		}
		c++
	}
}
