package main

import "fmt"
import "jumper/common/flexi"

func main() {


    var myIntArray    = []int {}
    var myStringArray = []string{"a","b"}

    var myInterfaceSlice = make([]interface{},0)
    //var mySourceInterfaceSlice = make([]interface{},0)

    myInterfaceSlice = append(myInterfaceSlice, "a")
    myInterfaceSlice = append(myInterfaceSlice, 1)
    //
    var cSymbol       = "c"
    // var cSymboleSlice = []string{"c"}
    //myInterfaceSlice[] 


    newArray1,err1 := flexi.AppendInterface(myIntArray, cSymbol)
    newArray2,err2 := flexi.AppendInterface(myStringArray, cSymbol)
    newArray3,err3 := flexi.AppendInterface(myInterfaceSlice, cSymbol)

    newArray4,err4 := flexi.AppendInterfaceFrom(myInterfaceSlice, myInterfaceSlice)
    newArray5,err5 := flexi.AppendInterfaceFrom(myInterfaceSlice, cSymbol)
    newArray6,err6 := flexi.AppendInterface(myInterfaceSlice, myInterfaceSlice)


    fmt.Printf("1 test: |%v|%v|\n", newArray1,err1)
    fmt.Printf("2 test: |%v|%v|\n", newArray2,err2)
    fmt.Printf("3 test: |%v|%v|\n", newArray3,err3)

    fmt.Printf("4 test: |%v|%v|\n", newArray4,err4)
    fmt.Printf("5 test: |%v|%v|\n", newArray5,err5)
    fmt.Printf("6 test: |%v|%v|\n", newArray6,err6)

}
