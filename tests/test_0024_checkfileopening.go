package main

import "fmt"
import "os"

func main() {

    file, err := os.Open("/tmp/test")

    defer file.Close()

    if err != nil {

        return

    }

    file_info , err := file.Stat()

    if err != nil {

        return

    }

    size :=  file_info.Size()
    fmt.Printf("\n File size %d \n", size)




}
