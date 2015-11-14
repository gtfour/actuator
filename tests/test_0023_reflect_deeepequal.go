package main

import "reflect"
import "fmt"

type Test1 struct {

    Num     int
    String  string

}


func main() {

    test1:=&Test1{Num:2,String:"hello"}
    test2:=&Test1{Num:3,String:"hello"}
    test3:=&Test1{Num:2,String:"hello"}

    fmt.Printf("\ntest1 and test2 are equal : %t \n", reflect.DeepEqual(test1,test2))
    fmt.Printf("\ntest1 and test3 are equal : %t \n", reflect.DeepEqual(test1,test3))


}
