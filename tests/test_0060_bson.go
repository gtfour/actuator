package main

import "fmt"
import "gopkg.in/mgo.v2/bson"

func main() {

    test:=bson.M{"id": "hello"}
    fmt.Printf("\n%v\n",test)


}
