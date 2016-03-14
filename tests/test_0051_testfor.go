package main

import "fmt"

func main() {
    var test = []string {"0","1","2","3","4","5","6","7","8","9"}
    for i := range test {
        i = 4
        fmt.Printf("\n%s",test[i])
    }
}
