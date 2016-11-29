package main

import "os"
import "fmt"

func main(){

    fmt.Printf("Input args:\n%v\n",os.Args)
    for i:= range os.Args {
        arg:=os.Args[i]
        fmt.Printf("\n---%s---",arg)


    }
}
