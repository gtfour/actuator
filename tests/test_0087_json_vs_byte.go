package main

import "fmt"
import "encoding/json"

func main() {
    //
    //
    test           := make(map[string]interface{}, 0)
    test["Id"]     =  "0123"
    testStr        := fmt.Sprintf( "%v", test)
    testStrByte    := []byte(testStr)
    testJsByte,err := json.Marshal(test)
    fmt.Printf("\ntestStrByte: %v\ttestJsByte: %v\tjsErr: %v\t", testStrByte, testJsByte, err)
    //
    // byte repressentation
    //
}
