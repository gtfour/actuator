package main

import "fmt"
import "encoding/json"

func main(){

    my_map:=make(map[string]interface{},0)
    my_map_reciever:=make(map[string]interface{},0)
    my_map["hello"]="hello"
    my_map["я"]="ya"
    my_map["2"]="хуй"
    my_byte,err:=json.Marshal(my_map)
    fmt.Printf("\n%v\n%v",my_byte,err)
    err=json.Unmarshal(my_byte, &my_map_reciever)
    fmt.Printf("\nReciever:\n%v\n",my_map_reciever)
    for i,x := range my_map {
        fmt.Printf("%v  %v\n",i,x)
    }
}
