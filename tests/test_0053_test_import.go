package main

import . "wapour/api/webclient"
import . "wapour/api/webclient/userstorage"
import "fmt"


func main(){

    w:=WengineWrapper{UserId:"a",TokenId:"b",SessionId:"22"}
    test:=CreateUserStorage()
    if test!=nil {
        fmt.Printf("FindWrapper: %v\n",test.FindWrapper("a","b","22"))
        fmt.Printf("AddWrapper:  %v\n",test.AddWrapper(&w))
        fmt.Printf("FindWrapper: %v\n",test.FindWrapper("a","b","22"))
        fmt.Printf("\nUserStorage %v\n",test)
    }

}

