package main

import "fmt"
import "jumper/cuda"

func main(){

    target_config          := make(map[string]string,0)
    target_config["type"]  =  "SINGLE_LINE"
    tgt,err                := cuda.NewTarget(target_config)
    line                   := []string{"root:x:0:0:root:/root:/bin/bash"}
    tgt.AddLine(line)
    //
    fmt.Printf("\nTarget:\n%v\nError:\n%v\n",tgt,err)




}
