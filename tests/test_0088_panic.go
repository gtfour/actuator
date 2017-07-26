package main

import "fmt"

func main() {
    defer func() {
        //
        // catching panic
        if r := recover(); r != nil {
            fmt.Printf("%v",r)
            fmt.Printf("\n<= panic has been catched =>\n")
        }
        //
        //
    }()
    panic("\n---\n::a problem::\n---\n")
    fmt.Printf("continue after panic")
}
