package main

import "fmt"
import "client/cross"
import "client/common"

func main() {

    fmt.Printf("\n Storage error %v\n",cross.STORAGE_INSTANCE.Error)
    new_id,err:=common.GenId()
    dynima:=cross.Dynima{Id:"CB775B5F-F9A3-2CAA-7E26-87282824F7E1",SourcePath:"/etc/passwd",SourceType:"file"}
    cross.EditDynima(dynima)
    if err == nil {
        _ = cross.CreateDynima(new_id)
    }

    d,err:=cross.GetDynima("CB775B5F-F9A3-2CAA-7E26-87282824F7E1")
    fmt.Printf("\nDynima:\n%v\nErr:%v",d,err)
}
