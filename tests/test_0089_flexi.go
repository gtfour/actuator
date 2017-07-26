package main

import "fmt"
import "jumper/common/flexi"

func main() {


    var myIntArray    = []int {}
    var myStringArray = []string{"a","b"}
    var cSymbol       = "c"

    var myInterfaceSlice = make([]interface{},0)

    newArray1,err1 := flexi.AppendString(myIntArray, cSymbol)
    newArray2,err2 := flexi.AppendString(myStringArray, cSymbol)
    newArray3,err3 := flexi.AppendString(myInterfaceSlice, cSymbol)

    fmt.Printf("1 test: |%v|%v|\n", newArray1,err1)
    fmt.Printf("2 test: |%v|%v|\n", newArray2,err2)
    fmt.Printf("3 test: |%v|%v|\n", newArray3,err3)

}
