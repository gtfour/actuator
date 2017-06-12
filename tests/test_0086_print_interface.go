package main

import "fmt"

func main() {

    myMap  := make(map[string]interface{}, 0)
    myMap["value1"] = "my_string"
    myMap["value2"] = 2
    //
    myNewStr1 := fmt.Sprintf("%v", myMap["value1"])
    myNewStr2 := fmt.Sprintf("%v", myMap["value2"])
    fmt.Printf("\nConverting interface to string: %s %s\n", myNewStr1, myNewStr2)
    //
    myMap["value3"] = []byte(myNewStr1)
    myMap["value4"] = []byte(myNewStr2)

    fmt.Printf("Printing whole map:\n%v\n", myMap)



}
