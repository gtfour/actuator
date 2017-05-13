package main

import "fmt"

var HUI int = 1000

func main() {
    fmt.Printf("\n---\nTest result: %v\n---\n", Test())
}

func Test()(int) {
    mySlice := []int{ 88, 90, 92, 94, 96}
    //
    for i := range mySlice {
        elem:=mySlice[i]
            if elem == 92 { return  elem }
    }
    //
    return HUI
}
