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
    // directory target
    targetDirectoryConfig              := make(map[string]string,0)
    targetDirectoryConfig["type"]      = "TARGET_DIRECTORY"
    targetDirectoryConfig["path"]      =  "/etc/wengine/"
    tgtDirectory,err                   := cuda.NewTarget(targetDirectoryConfig)
    fmt.Printf("\nTarget:\n%v\nError:\n%v\n---\n",tgtDirectory,err)
    err=tgtDirectory.Gather()
    directory_nested_targets:=tgtDirectory.GetNestedTargets()
    fmt.Printf("%v\n---\nNested Targets:%v\nError:%v\n-----", tgtDirectory, directory_nested_targets, err)
    // --
    for i:= range directory_nested_targets {
        myNestedTarget:=directory_nested_targets[i]
        fmt.Printf("\n%v",*myNestedTarget)
    }
    fmt.Printf("\n-----\n")
}
