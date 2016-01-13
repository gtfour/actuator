package main

import "fmt"
//
import "client/actuator"

func main() {

    path     := "/tmp/test"
    dir      := &actuator.Directory{}
    err      := dir.GetHashSumDir(path)
    if err== nil {

        for i:=range dir.Files {
            file:=dir.Files[i]
            fmt.Printf("\n --- \n")
            fmt.Println(file.Path)
            fmt.Println(string(file.Prop.HashSum))
            fmt.Printf("\n --- \n")


        }


    }


}
