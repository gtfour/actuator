package main

import "fmt"
//import "gopkg.in/mgo.v2/bson"

func Test( test ...map[string]string) (hello string) {


    _,ok:=test[0]["hello"]
    fmt.Printf("--\n%v\n", ok)
    return "hello"


}

func main() {

    //test:=make(map[string]interface{})
    test:=make(map[string]string)
    test["Name"] = "John"
    test["SecondName"] = "Johnson"
    test["hello"]="sdcsd"
    //Test(test, test)
    //fmt.Printf("===\n%v\n===",bson.M(test))
    Test(test)
}
