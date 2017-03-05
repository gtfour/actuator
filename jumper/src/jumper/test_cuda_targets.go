package main

import "fmt"
import "jumper/cuda"

func main(){
    //
    // -- 
    // line target
    target_config          := make(map[string]string,0)
    target_config["type"]  =  "SINGLE_LINE"
    tgt,err                := cuda.NewTarget(target_config)
    line                   := []string{"root:x:0:0:root:/root:/bin/bash"}
    tgt.AddLine(line)
    //
    fmt.Printf("\nTarget:\n%v\nError:\n%v\n",tgt,err)
    // --
    // file target
    targetFileConfig           := make(map[string]string,0)
    targetFileConfig["type"]   =  "TARGET_FILE"
    targetFileConfig["path"]   =  "/etc/wengine/wengine.conf"
    tgtFile,err                := cuda.NewTarget(targetFileConfig)
    fmt.Printf("\nTarget:\n%v\nError:\n%v\n---\n",tgtFile,err)
    err=tgtFile.Gather()
    fmt.Printf("%v\n---\nError:%v\n", tgtFile,err)
    // --
    //
}
