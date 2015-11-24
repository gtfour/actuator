package main

import "client_side/sistory"
import "fmt"


func main() {

storage:=sistory.Open()

fmt.Printf("Error:  %v",storage.Error)

storage.Close()




}
