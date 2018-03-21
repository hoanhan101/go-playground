package main

import (
    "fmt"
    "time"
)

type Ball struct {
    hits int
}

func main() {
    table := make(chan *Ball)

    go player("ping", table)
    go player("pong", table)

    // Game on, start the ball
    table <- new(Ball)

    time.Sleep(1 * time.Second)

    // Gameover, grab the ball
    <-table

    // This will show us the goroutine is still running
    // because we are in the infinity loop. Very wasteful.
    panic("Show me the stacks")
}

func player(name string, table chan *Ball) {
    for {
        ball := <-table
        ball.hits++
        fmt.Println(name, ball.hits)
        time.Sleep(100 * time.Millisecond)
        table <- ball
    }
}
