package main

import "fmt"
import "gopkg.in/mgo.v2/bson"

func Test( test ...map[string]string) (hello string) {

    fmt.Printf("--\n%v\n--", test)
    return "hello"


}

func main() {

    test:=make(map[string]interface{})
    test["Name"] = "John"
    test["SecondName"] = "Johnson"
    //Test(test, test)
    fmt.Printf("===\n%v\n===",bson.M(test))
}
