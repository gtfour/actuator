package main

import "fmt"
import "jumper/common/file"


func main(){

    lines,offset,err:=file.ReadFileWithOffset("/etc/wengine/wengine.conf",14)
    fmt.Printf("Lines:%v\nOffset:%v\nErr:%v",lines,offset,err)


}
