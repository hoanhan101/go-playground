package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()

	fmt.Println("main")
	fmt.Println("NumCPU: ", runtime.NumCPU())
	fmt.Println("NumCgoCall: ", runtime.NumCgoCall())
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
}

func foo() {
	fmt.Println("Foo")
	fmt.Println("NumCPU: ", runtime.NumCPU())
	fmt.Println("NumCgoCall: ", runtime.NumCgoCall())
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())

	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	fmt.Println("Bar")
	fmt.Println("NumCPU: ", runtime.NumCPU())
	fmt.Println("NumCgoCall: ", runtime.NumCgoCall())
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())

	for i := 0; i < 10; i++ {
		fmt.Println("bar: ", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}
