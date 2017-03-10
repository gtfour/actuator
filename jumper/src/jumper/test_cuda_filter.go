package main

import "fmt"
import "jumper/cuda"

func main(){

    var filterList2 FilterList
    filterList1:=cuda.CreateDefaultFilterList()
    fmt.Printf("<<<\nFilter list: %v\n%v\n>>>", filterList1,len(filterList1))


}
