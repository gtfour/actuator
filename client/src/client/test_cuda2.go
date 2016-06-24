package main
import "fmt"
import "strings"
import "client/blacktop"
import "client/cuda"


//import "fmt"
//import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {


    lines,err:=blacktop.ReadFileLines("/tmp/hello.txt")
    if (err==nil) {


    for i:= range lines {
        line:=lines[i]
        lineAsArray3:=strings.Split(line, "")
        delims,data:=cuda.GetIndexes(lineAsArray3)
        //delims,data=cuda.PathFilter(lineAsArray3,delims,data)
        delims,data=cuda.QuotesFilter(lineAsArray3,delims,data)
        delims,data=cuda.PathFilter(lineAsArray3,delims,data)
        delims,data=cuda.UrlFilter(lineAsArray3,delims,data)
        template_string:=cuda.GenTemplate(lineAsArray3, data)
        fmt.Printf("line:%s\n",line)
        fmt.Printf("Template string:%s\n",template_string)

    }

    }






}
