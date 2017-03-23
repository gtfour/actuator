package main

import "fmt"
import "jumper/cuda/targets"

func main(){
    //
    // -- 
    // line target
    target_config          := make(map[string]string,0)
    target_config["type"]  =  "SINGLE_LINE"
    tgt,err                := targets.NewTarget(target_config)
    line                   := []string{"root:x:0:0:root:/root:/bin/bash"}
    tgt.AddLine(line)
    //
    fmt.Printf("\n>>    Line Target:    <<\n%v\nError:\n%v\n>>        <<",tgt,err)
    // -- -- -- --
    // file target
    // -- -- -- --
    targetFileConfig           := make(map[string]string,0)
    targetFileConfig["type"]   =  "TARGET_FILE"
    targetFileConfig["path"]   =  "/etc/wengine/wengine.conf"
    tgtFile,err                := targets.NewTarget(targetFileConfig)
    fmt.Printf("\n>>    File Target:    <<\n%v\nError:\n%v\n---\n",tgtFile,err)
    err=tgtFile.Gather()
    fmt.Printf("%v\n---\nError:%v\n>>        <<", tgtFile,err)
    // -- -- -- --
    // directory target
    // -- -- -- --
    targetDirectoryConfig              := make(map[string]string,0)
    targetDirectoryConfig["type"]      = "TARGET_DIRECTORY"
    targetDirectoryConfig["path"]      =  "/etc/wengine/"
    tgtDirectory,err                   := targets.NewTarget(targetDirectoryConfig)
    fmt.Printf("\n>>    Directory Target:    <<\n%v\nError:\n%v\n---\n",tgtDirectory,err)
    err=tgtDirectory.Gather()
    directory_nested_targets           := tgtDirectory.GetNestedTargets()
    fmt.Printf("%v\n---\n>>    Nested Targets:    <<\n%v\nError:%v\n-----", tgtDirectory, directory_nested_targets, err)
    //
    // -- -- --
    //
    for i:= range directory_nested_targets {
        myNestedTarget:=directory_nested_targets[i]
        fmt.Printf("\n\t%v\n",*myNestedTarget)
    }
    fmt.Printf("\n>>        <<\n")
    //
    // -- -- --
    //
}
