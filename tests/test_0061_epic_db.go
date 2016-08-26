package main

import "fmt"
import "gopkg.in/mgo.v2/bson"

func main() {

    my_map:=make(map[string]interface{},0)
    my_map["id"]="222"
    my_bson_map:=bson.M(my_map)

    test:=bson.M{"id": "hello"}
    fmt.Printf("\n%v\n",test)
    fmt.Printf("\nNew Bson Map:%v",my_bson_map)


}
