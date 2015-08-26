package main

import "fmt"

func main() {


StringSearch("hello","h")




}

func StringSearch(line string, char string) {

    for i := range line {

        state:=char==string(line[i])

        fmt.Printf("%s %d %t\n",string(line[i]),line[i],state)



    }




}
