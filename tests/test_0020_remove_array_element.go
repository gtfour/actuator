package main

import "fmt"

func main(){

var a = []int {1,2,3,4,5,6}


a = append(a[:0], a[0+1:]...)

for i := range a {

    fmt.Printf("%d",a[i])


}
}
