package main

import "fmt"

type Searcher struct {
    // //
    // //
    value               string
    since               int
    direction           int
    maxCount            int
    // //
    // //  should satisfy to Accepter and Breaker . if  Accepter returns true and Breaker return false searching will remain
    accepter            func(string)(bool)
    breaker             func(string)(bool)
    resultPosition      int // calculating field
    // //
    // //
}



func main() {

    searcher := Searcher{}
    if searcher.breaker == nil {
        fmt.Printf("\nbreaker is nil\n")
    }



}
