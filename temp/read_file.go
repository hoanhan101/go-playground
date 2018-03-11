package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    content, err := ioutil.ReadFile("temp.txt")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Before: %s %T \n", content, content)
    fmt.Printf("After: %s %T \n", string(content), string(content))
}
