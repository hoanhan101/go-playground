package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(1 * time.Second)

	go func() {
		fmt.Println(time.Now().Sub(start))
	}()

	go func() {
		fmt.Println(time.Since(start))
	}()

	time.Sleep(1 * time.Second)
}
