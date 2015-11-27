package main

import "client_side/sistory"
import "fmt"
import "encoding/json"



func main() {

    var test map[string]interface{} // kpmý@golang.cjr advice

    storage:=sistory.Open()

    spirit_prop:=&sistory.SpiritProp{Size:122,Path:"/tmp/test/hello22.txt"}

    err:=storage.UploadSpirit(spirit_prop)

    fmt.Printf("%v",err)

    data:=storage.CallSpirit("/tmp/test/hello22.txt")

    fmt.Printf("Data: %s\n",string(data))
    err = json.Unmarshal(data, &test)

    fmt.Printf("%v %v\n", test, err)
    fmt.Printf("%s\n", test["192.168.236.11"])

    storage.Close()

}
