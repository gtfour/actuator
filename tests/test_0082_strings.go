package main

import "fmt"
import "strings"

func main(){

    //hello1 := "[base]"
    hello2 := "base"

    oindex := strings.Index(hello2,"[")
    cindex := strings.Index(hello2,"]")
    fmt.Printf("\nindexes: %d  %d\n",oindex,cindex)


}
