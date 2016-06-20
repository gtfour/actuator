package main

import "fmt"
import "client/cross"
import "client/common"

func main() {

    fmt.Printf("\n Storage error %v\n",cross.STORAGE_INSTANCE.Error)
    new_id,err:=common.GenId()

    var data =  [][]string {[]string{"a"},[]string{"b"},[]string{"c"}}

    dynima:=cross.Dynima{Id:new_id,SourcePath:"/etc/passwd",SourceType:"file", Data:data}
    dynima.Write()
    //if err == nil {
    //    _ = cross.CreateDynima(new_id)
    //}

    d,err:=cross.GetDynima(new_id)
    fmt.Printf("\nDynima:\n%v\nErr:%v",d,err)
    dns,err:=cross.GetDynimasByPath("/etc/passwd")
    for i:= range dns {
        fmt.Printf("--\n%v\n--",dns[i])
    }
}
