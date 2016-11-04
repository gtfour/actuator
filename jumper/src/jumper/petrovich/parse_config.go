package petrovich

import "fmt"
import "jumper/common/file"

func ParseConfig(filename string)(file_config map[string]string,err error){
    //
    //
    //
    lines, err := file.ReadFile(filename)
    fmt.Printf("-- hello --\n")
    fmt.Printf("%v",lines)
    fmt.Printf("\n-- --")
    return file_config, err
    //
    //
    //
}
