package main

import "fmt"
import "encoding/json"

func main(){

    my_map          := make(map[string]interface{},0)
    my_map["hello"]="hello"
    my_map["я"]="ya"
    my_map["2"]="хуй"

    my_map_reciever := make(map[string]interface{},0)

    my_string       := "hello"

    var my_string_reciever string

    fmt.Printf("\n Decoder: \n")

    my_map_byte,err_mb    := json.Marshal(my_map)
    my_string_byte,err_sb := json.Marshal(my_string)

    fmt.Printf("\nMap:    %v   %v\n",my_map_byte,err_mb)
    fmt.Printf("\nString: %v   %v\n",my_string_byte,err_sb)

    fmt.Printf("\n Encoder: \n")

    err_map_to_map    := json.Unmarshal(my_map_byte, &my_map_reciever)
    err_map_to_string := json.Unmarshal(my_map_byte, &my_string_reciever)

    fmt.Printf("\nMap to Map Reciever:\n%v\nErr:%v",my_map_reciever,err_map_to_map)
    fmt.Printf("\nMap to String Reciever:\n%v\nErr:%v",my_string_reciever,err_map_to_string)

    err_string_to_map := json.Unmarshal(my_string_byte, &my_map_reciever)

    fmt.Printf("\nString to Map Reciever:\n%v\nErr:%v",my_map_reciever,err_string_to_map)

    //for i,x := range my_map {
    //    fmt.Printf("%v  %v\n",i,x)
    //}
}
