package main

import "fmt"
import "strings"

func main() {


//StringSearch("hello","h")

string_without_equal_sign:=`Hello "   hello h`

fmt.Println(strings.Replace(string_without_equal_sign, `"`, "-", -1))

fmt.Printf(" Len %d\n",len(strings.Split(string_without_equal_sign,"=")))


}

func StringSearch(line string, char string) {

    for i := range line {

        state:=char==string(line[i])

        fmt.Printf("%s %d %t\n",string(line[i]),line[i],state)



    }




}
