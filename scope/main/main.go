package main

import (
    "fmt"
    "github.com/hoanhan101/go-playground/scope/visible"
)

func main() {
    // Can use MyName because is uppercase, visible outside the package
    // Cannot use myName because it is not visible outside the package
    // However, if we call PrintMyName, which is visible outside the package
    // and in PrintMyName, it can print myName because they are in the same
    // package scope level
    fmt.Println(visible.MyName)
    visible.PrintMyName()
}
