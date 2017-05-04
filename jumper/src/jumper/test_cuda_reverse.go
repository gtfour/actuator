package main

import "fmt"
import "jumper/cuda/analyze"

func main() {

    //myString := "::1:b22:c33:d44:e55:::::::::::::::"
    myString := "::2::"
    sliceLen := len(myString)
    // data_indexes :=[][]int{[]int{2,2}, []int{4,6}, []int{8,10}, []int{12,14}, []int{16,18}}
    data_indexes :=[][]int{ []int{2,2} }
    reverse:=analyze.MakeReverse(data_indexes, sliceLen)
    fmt.Printf("\nOrig:    %v\n", data_indexes)
    fmt.Printf("\nReverse: %v\n", reverse)


}
