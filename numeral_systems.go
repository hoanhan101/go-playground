package main

import "fmt"

func main() {
	number := 101
	fmt.Printf("Decimal represenation of number %d is %d\n", number, number)
	fmt.Printf("Binary represenation of number %d is %b\n", number, number)
	fmt.Printf("Hexadecimal represenation of number %d is %x\n", number, number)
	fmt.Printf("Hexadecimal represenation of number %d with leading 0 is %#x\n", number, number)

	for i := 0; i < 10; i++ {
		fmt.Printf("Decimal: %d \t Binary: %b \t Hex: %#x \t UTF-8: %q \n", i, i, i, i)
	}
}
