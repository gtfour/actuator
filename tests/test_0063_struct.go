package main

import "fmt"

type Query struct {
    Type      int
    Table     string
    KeyBody   map[string]interface{}
    QueryBody map[string]interface{}
}


func main(){

    my_query:=Query{KeyBody:"hello"}
    fmt.Printf("\n%v\n",my_query)


}
