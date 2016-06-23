package main

import "fmt"
import "client/blacktop"

func main() {

str,err:=blacktop.ReadFileLines("/etc/passwd")
fmt.Printf("\n%v\n",err)

for i:=range str {
    fmt.Printf("[%d]",i)
    fmt.Println(str[i])
}


}
