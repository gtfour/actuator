package main

import "fmt"
import "jumper/cuda"

func main(){
    //
    var filterList2   cuda.FilterList
    var myFilter      cuda.Filter
    myFilter.Enabled  = true
    myFilter.Call     = filter
    myFilter.Name     = "myFilter"
    //
    filterList1       := cuda.CreateDefaultFilterList()
    //
    //
    //
    err_on_append     := filterList2.Append(myFilter)
    //
    fmt.Printf("\nAppend error is %v\n", err_on_append)
    //
    fmt.Printf("<<<\nFilter list #1: %v\n%v\n>>>", filterList1, len(filterList1))
    fmt.Printf("<<<\nFilter list #2: %v\n%v\n>>>", filterList2, len(filterList2))

    //
    //
}


func filter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int){
    return
}
