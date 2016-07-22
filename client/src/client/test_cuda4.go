package main
import "client/cuda"
import "fmt"
import "strings"

func main() {

    line3:=`"hello":"Jessie:'http:///www.google.com'"`

    cuda.DebugPrintCharCounter(line3)

    lineAsArray3:=strings.Split(line3, "")

    delims,data:=cuda.GetIndexes(lineAsArray3)

    fmt.Printf("\nBefore: delims: %v\n  data: %v \n" , delims , data)

    delims,data=cuda.QuotesFilter(lineAsArray3,delims,data)

    fmt.Printf("\nAfter QuotesFilter:  delims: %v\n data: %v \n" , delims , data)

    delims,data=cuda.PathFilter(lineAsArray3,delims,data)

    fmt.Printf("\nAfter PathFilter:  delims: %v\n data: %v \n" , delims , data)

    delims,data=cuda.UrlFilter(lineAsArray3,delims,data)

    fmt.Printf("\nAfter UrlFilter:  delims: %v\n data: %v \n" , delims , data)

    template_string:=cuda.GenTemplate(lineAsArray3, data)

    fmt.Printf("###\nTemplate string:\n%s\n",template_string)

}

