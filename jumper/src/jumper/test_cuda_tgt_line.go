package main

import "fmt"
import "jumper/cuda/targets"
import "jumper/cuda/handling"
import "jumper/cuda/filtering"


func main(){
    lines := []string{"root:x:0:0:root:/root:/bin/bash","vagrant:x:1002:1002::/home/vagrant:","www-data:x:33:33:www-data:/var/www:/usr/sbin/nologin"}
    handler :=  handling.NewHandler(nil)
    fl      :=  filtering.CreateDefaultFilterList()
    handler.AddFilters(fl)
    target_config            := make(map[string]string, 0)
    target_config["type"]    =  "SINGLE_LINE"
    tgt,_                  := targets.NewTarget(target_config)
    handler.AddTargetPtr(tgt)
    for i := range lines {
        line:=lines[i]
        tgt.SetLine(line)
        result,err :=  handler.Handle()
        fmt.Printf("result: %v err: %v",result,err)
    }
}
