package main

import "client/sisb"
import "client/common"

func main() {

    new_id,err:=common.GenId()
    if err == nil {
        _ = sisb.CreateDynima(new_id)
    }



}

