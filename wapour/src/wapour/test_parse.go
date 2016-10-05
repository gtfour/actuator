package main

import "fmt"
import "wapour/core/parse"

func main(){
    lines,err:=parse.ReadFileLines("/etc/wapour/settings.go")
    fmt.Printf("error:%v\nlines:\n%v",err,lines)
}
