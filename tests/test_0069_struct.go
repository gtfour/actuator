package main

import "fmt"
import "encoding/json"

type Test struct {
    KeyBody     map[string]interface{}
    MyTest
}

type MyTest struct {
    QueryBody   map[string]interface{}
}

func main(){
    x:=Test{}
    if x.KeyBody == nil {
        fmt.Printf("\nx.KeyBody is nil\n")
    }
    mybyte,err:=json.Marshal(x.KeyBody)
    fmt.Printf("\nEncoded byte %v | %v",mybyte,err)


}
