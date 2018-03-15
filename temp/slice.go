package main

import (
	"fmt"
)

func main() {
	var buffer [256]byte
	slice := buffer[10:20]

	// This change both the value under slice and buffer because a slice is
	// holds pointers to elements in buffer
	slice[0] = 1
	fmt.Println(slice)
	fmt.Println(buffer[10:20])
}
