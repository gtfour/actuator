package main
import "fmt"
import "strings"
import "client/blacktop"

import "jumper/cuda/analyze"
import "jumper/cuda/filtering"
import "jumper/cuda/templating"



//import "fmt"
//import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {


    lines,err:=blacktop.ReadFileLines("/tmp/hello.txt")
    if (err==nil) {


    for i:= range lines {
        line            := lines[i]
        lineAsArray3    := strings.Split(line, "")
        delims,data     := analyze.GetIndexes(lineAsArray3)
        //delims,data=cuda.PathFilter(lineAsArray3,delims,data)
        delims,data     =  filtering.QuotesFilter(lineAsArray3,delims,data)
        delims,data     =  filtering.PathFilter(lineAsArray3,delims,data)
        delims,data     =  filtering.UrlFilter(lineAsArray3,delims,data)
        template_string := templating.GenTemplate(lineAsArray3, data)
        fmt.Printf("line:%s\n",line)
        fmt.Printf("Template string:%s\n",template_string)

    }

    }






}
