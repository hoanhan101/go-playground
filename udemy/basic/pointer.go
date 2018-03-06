package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i            // point to i
	fmt.Println(i)     // value of i
	fmt.Println(&i)    // address of i
	fmt.Println(p)     // p is a pointer to i, so it holds i's address
	fmt.Println(&p)    // & of an address is another address
	fmt.Println(*p)    // * of an address is a value
	fmt.Println(*(&i)) // this case, &i is an address, so *(&i) is a value

	// because p holds the address of i, *p is i, which is 42
	// this will set the value of i through p
	*p = 21
	fmt.Println(i) // see the new value of i

	fmt.Println("new chapter") // see the new value of i
	p = &j                     // point to j
	fmt.Println(p)             // see the new value of i

	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
