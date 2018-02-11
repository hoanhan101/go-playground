package visible

import (
    "fmt"
)

func PrintMyName() {
    fmt.Println("Printing my name from visible package")
    fmt.Println("MyName:", MyName)
    fmt.Println("myName:", myName)
}
