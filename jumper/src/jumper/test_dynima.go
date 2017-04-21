package main


import "fmt"
import "jumper/cuda"
import "jumper/cuda/targets"
import "jumper/cuda/filtering"

func main(){
    // -- -- -- --
    // directory target
    // -- -- -- --
    targetDirectoryConfig              := make(map[string]string,0)
    targetDirectoryConfig["type"]      = "TARGET_DIRECTORY"
    targetDirectoryConfig["path"]      =  "/etc/wengine/"
    tgtDirectory,err                   := targets.NewTarget(targetDirectoryConfig)
    if err!=nil { fmt.Printf("\n Directory config error: %v \n", err)  }
    err=tgtDirectory.Gather()
    if err!=nil { fmt.Printf("\n Directory gather error: %v \n", err)  }
    //
    // -- -- --
    //
    //
    // -- -- -- --
    // dynima to handle directory target
    // -- -- -- --
    //
    d                 := cuda.Dynima{}
    defaultFilterList := filtering.CreateDefaultFilterList()
    for i:= range filtering.CreateDefaultFilterList(){
        filter := defaultFilterList[i]
        d.AppendFilter(filter)
    }
    d.AppendTarget(tgtDirectory)
    fmt.Printf("\n:Dynima: %v\n", d )
    result := d.RunFilters()
    fmt.Printf("\n:Result: %v\n", result )
    //
    //
    //
}
