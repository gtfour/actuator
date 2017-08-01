package main

import "fmt"
import "jumper/common/flexi"

func main() {
    //
    //
    var err error = nil
    inSlice1:=make([]interface{},0)
    //
    //
    inSlice1,err=flexi.Append(inSlice1,"e")
    inSlice1,err=flexi.Append(inSlice1,4)
    inSlice1,err=flexi.Append(inSlice1,false)
    inSlice1,err=flexi.Append(inSlice1,"zz")
    inSlice1,err=flexi.Append(inSlice1,10010912929)
    inSlice1,err=flexi.Append(inSlice1,true)
    inSlice1,err=flexi.Append(inSlice1,[]int{1,2,3,4})
    inSlice1,err=flexi.Append(inSlice1,false)
    //
    //
    removeIndexes := []int{2,5}
    fmt.Printf("\nslice : %v \nnow we will remove following indexes :%v  from this slice", inSlice1, removeIndexes)
    inSlice1,err=flexi.Remove(inSlice1, removeIndexes) // remove bool elems // fun bug :)))
    fmt.Printf("\nResult is: %v\nerr:%v\n", inSlice1,err)
    singleRemoveIndex := len(inSlice1)-1
    fmt.Printf("\nAnd now we just remove single one(last one): %v\n", singleRemoveIndex)
    inSlice1,err=flexi.Remove( inSlice1, singleRemoveIndex )
    fmt.Printf("\nAnd result is: %v\nerr:%v\n",inSlice1,err)
    //
    // fmt.Printf("\ninSlice1: %v\nerr:%v", inSlice1,err)
}
