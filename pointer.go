package main

import "fmt"

func main() {
	// a is an integer
	a := 101

	// Print an integer a
	fmt.Println(a)

	// Print the address of a
	fmt.Println(&a)

	// b is a pointer to an integer
	var b *int = &a

	// b is the address of a
	fmt.Println(b)

	// Dereference b, which is the value a
	fmt.Println(*b)

	// The address of the address b is another address
	fmt.Println(&b)

	// Change the value at address b to something else
	*b = 202
	fmt.Println(*b)

	// Because address b is is a
	// or *b = a
	// changing *b also changes a
	fmt.Println(a)

	// Demo pointer
	// Instead of passing the value c, we're passing the address of c
	c := 1
	demoPointer(&c)
	fmt.Println(c)
}

// The parameter is a pointer to an integer
func demoPointer(target *int) {
	// Defeference that address of the target and increase the value at the address
	*target += 100
}
