package main

import (
    "os"
    "fmt"
    "log"
)

func init() {
    nf, err := os.Create("log.txt")
    if err != nil {
        fmt.Println(err)
    }
    log.SetOutput(nf)
}

func main() {
    _, err := os.Open("ghost.txt")
    if err != nil {
        // Way 1
        // fmt.Println(err)

        // Way 2
        // log.Println(err)

        // Way 3
        // Call os.Exit() after writing the log message
        // log.Fatalln(err)

        // Way 4
        panic(err)
    }
}
