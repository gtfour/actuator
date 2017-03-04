package main

import "fmt"
import "jumper/common/file"

func main(){
    dir_name:="/etc/wengine"
    dir_content,_:=file.ReadDir(dir_name)
    fmt.Printf("Dir %s Content:\n%v\n", dir_name, dir_content)
}
