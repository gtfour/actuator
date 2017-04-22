package main

import "fmt"
import "jumper/cuda/analyze"

func main() {

    //
    pairCase11  := [2]int{0,0}
    pairCase12  := [2]int{5,5}
    //
    pairCase21  := [2]int {0,2}
    pairCase22  := [2]int {4,6}
    pairCase23  := [2]int {8,13}
    pairCase24  := [2]int {15,15}


    //
    indexesCase1 := [][2]int {pairCase11,pairCase12}
    indexesCase2 := [][2]int {pairCase21, pairCase22, pairCase23, pairCase24}
    //
    lineAsArrayCase1 := []string {"a"," ","="," ",`"`,"3",`"`}
    lineAsArrayCase2 := []string {"a","b","c",":","v","r","2",":","3","2","2","2","4","4",":","b"}
    //
    res1 := analyze.SelectDataByIndexes(lineAsArrayCase1, indexesCase1)
    res2 := analyze.SelectDataByIndexes(lineAsArrayCase2, indexesCase2)
    //
    fmt.Printf("\nResult1:%v\n" , res1)
    fmt.Printf("\nResult2:%v\n" , res2)


}
