package main

import (
    "errors"
    "log"
    "fmt"
)

var ErrNegativeSqrt = errors.New("sqrt: square root of negative number")

func main() {
    _, err := Sqrt(-10)
    if err != nil {
        log.Fatalln(err)
    }
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        // Way 1
        // return 0, ErrNegativeSqrt

        // Way 2
        return 0, fmt.Errorf("sqrt: asquare toot of negative number: %v", f)
    }
    return 1, nil
}
