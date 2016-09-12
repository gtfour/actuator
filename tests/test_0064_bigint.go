package main

import (
    "fmt"
    "math/big"
)

func main() {
    verybig := big.NewInt(1)
    ten     := big.NewInt(10)
    for i:=0; i<100000; i++ {
       temp := new(big.Int)
       temp.Mul(verybig, ten)
       verybig = temp
    }
    //var simple_first int = 100000000000
    //first  := big.NewInt(100000000000)
    //fmt.Println(simple_first)
    //fmt.Println(first)
    fmt.Println(verybig)
}
