package main

import "fmt"

func main(){

var a = []int {1,2,3,4,5,6}


a = append(a[:1], a[1+1:]...)

for i := range a {

    fmt.Printf("%d",a[i])


}
}
