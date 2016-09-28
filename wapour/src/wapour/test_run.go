package main

import "fmt"
import "wapour/core/run"

func main(){

    settings,err:=run.GetCurrentAppSettings("/etc/wapour/settings.go")
    fmt.Printf("Settings:\n%v\n",settings,err)


}
