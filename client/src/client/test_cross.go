package main

import "fmt"
import "client/cross"
import "client/common"

func main() {

    fmt.Printf("\n Storage error %v\n",cross.STORAGE_INSTANCE.Error)
    new_id,err:=common.GenId()
    dynima:=cross.Dynima{Id:new_id,SourcePath:"/etc/passwd",SourceType:"file"}
    dynima.Write()
    //if err == nil {
    //    _ = cross.CreateDynima(new_id)
    //}

    d,err:=cross.GetDynima(new_id)
    fmt.Printf("\nDynima:\n%v\nErr:%v",d,err)
}
