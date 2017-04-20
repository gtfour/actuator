package main

import "fmt"

func main(){

    var indexes =  []int {1,2}
    var word    string = "abcdef"
    fmt.Printf("\n%v\n",word[indexes[0]:indexes[1]])


}
