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
    targetDirectoryConfig["path"]      = "/tmp/repos/nginx/"
    tgtDirectory,err                   := targets.NewTarget(targetDirectoryConfig)
    //
    if err!=nil { fmt.Printf("\n Directory config error: %v \n", err)  }
    //err=tgtDirectory.Gather()
    //if err!=nil { fmt.Printf("\n Directory gather error: %v \n", err)  }
    //
    // checking directory files
    //
    // nestedTargets:=tgtDirectory.GetNestedTargets()
    // for i:= range nestedTargets {
    //     target := nestedTargets[i]
    //    fmt.Printf("\nTarget: %v\n", target)
    // }
    //
    //-- dynima to handle directory target  --
    d                 := cuda.Dynima{}
    defaultFilterList := filtering.CreateDefaultFilterList()
    //
    //
    //
    for i := range filtering.CreateDefaultFilterList() {
        //
        //
        filter := defaultFilterList[i]
        d.AppendFilter( filter )
        //
        //
    }
    d.AppendTarget( tgtDirectory )
    fmt.Printf("\n:Dynima:%v\n-- -- -- --\n", d )
    resultSet := d.RunFilters()
    fmt.Printf("\n:ResultSet:\n-- -- -- --\n")
    results,_ := resultSet.GetData()
    //
    //
    //
    for i := range results {
        //
        // 
        result         := results[i]
        resultByte,err := result.GetJson()
        fmt.Printf("\n%s\nErr:\n%v",string(resultByte),err)
        //
        //
    }
    //
    //
    //
}
