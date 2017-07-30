package main

import "fmt"
import "jumper/common/flexi"

func main() {
    var err error = nil
    inSlice1:=make([]interface{},0)

    inSlice1,err=flexi.Append(inSlice1,"e")
    inSlice1,err=flexi.Append(inSlice1,4)
    inSlice1,err=flexi.Append(inSlice1,false)
    inSlice1,err=flexi.Append(inSlice1,"zz")
    inSlice1,err=flexi.Append(inSlice1,10010912929)
    inSlice1,err=flexi.Append(inSlice1,true)
    inSlice1,err=flexi.Remove(inSlice1,[]int{2,4}) // remove bool elems // fun bug :)))
    fmt.Printf("\ninSlice1: %v\nerr:%v", inSlice1,err)




}
