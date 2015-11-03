package main

import "fmt"

func main(){

    var test  = []int {1,2,4,5}
    for i:=range test {

        if i== 2 { test[i]=10 }

    }


    for i:=range test {

        fmt.Printf("%d",test[i])


    }



}
