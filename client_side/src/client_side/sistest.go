package main

import "client_side/sistory"
import "fmt"


func main() {

storage:=sistory.Open()

spirit_prop:=&sistory.SpiritProp{Size:122,Path:"/tmp/test/hello22.txt"}
_=storage.UploadSpirit(spirit_prop)
data:=storage.CallSpirit("/tmp/test/hello22.txt")


fmt.Printf("Data:  %s",string(data))

storage.Close()




}
