package main

import "fmt"
import "jumper/common/file"

func main(){
    dir_content,_:=file.ReadDir("/etc/wapour/")
    fmt.Printf("Dir Content:%v",dir_content)

}
