//
// Timeout using select 
// The time.After function returns a channel that blocks for the specified
// duration. After the interval, the channel delivers the current time, once.
//

package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
    c := boring("Hoanh")
    timeout := time.After(5 * time.Second)

    for {
        select {
        case s := <-c:
            fmt.Println(s)
        case <-timeout:
            fmt.Println("You talk to much")
            return
        default:
            fmt.Println("hmm")
        }
    }
}

func boring(msg string) <-chan string {
    c := make(chan string)

    go func() {
        for i := 0; ; i++ {
            c <- fmt.Sprintf("%s %d", msg, i)
            time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
        }
    }()

    return c
}
