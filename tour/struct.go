package main

import (
	"fmt"
	"strconv"
)

func main() {
	hoanh := Person{"Hoanh", "An", 21}
	hoanh.toString()
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) toString() {
	fmt.Printf("%v %v %v", p.firstName, p.lastName, strconv.Itoa(p.age))
}
