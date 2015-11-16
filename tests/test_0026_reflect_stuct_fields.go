package main

import "fmt"
import "reflect"

type StructOne struct {

    TestInt    int
    TestString string


}

type StructTwo struct {

    TestInt       int
    TestString    string
    TestIntSecong int


}

func main(){

    test1:=StructOne{TestInt:2,TestString:"Hello"}
    test2:=&StructTwo{}

    if  test2 == nil { fmt.Printf("\n:: second struct obj pointer is nil :: \n")}

    value1:=reflect.ValueOf(test1)
    value2:=reflect.ValueOf(test2)

    fmt.Printf("\nLen of StructOne  %d\n",value1.NumField())
    fmt.Printf("\nLen of StructTwo  %d\n",value2.NumField())


}
