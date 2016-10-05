package main

//import "fmt"
import "wapour/core/parse"

func main(){
    _,_=parse.ReadFileLines("/etc/wapour/wapour.conf")
    //fmt.Printf("error:%v\nlines:\n%v",err,lines)
}
