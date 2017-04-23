package main

import "jumper/cuda/analyze"
import "jumper/cuda/filtering"
import "fmt"
import "strings"

func main() {

    //line3:=`"hello":"Jessie:'http:///www.google.com'"`
    line3:="hello = https://translate.google.ru/#en/ru/duplicate"

    analyze.DebugPrintCharCounter(line3)

    lineAsArray3:=strings.Split(line3, "")

    delims,data:=analyze.GetIndexes(lineAsArray3)

    fmt.Printf("\nBefore: delims: %v\n  data: %v \n" , delims , data)

    delims,data=filtering.QuotesFilter(lineAsArray3,delims,data)

    fmt.Printf("\nAfter QuotesFilter:  delims: %v\n data: %v \n" , delims , data)

    delims,data=filtering.PathFilter(lineAsArray3,delims,data)

    fmt.Printf("\nAfter PathFilter:  delims: %v\n data: %v \n" , delims , data)

    delims,data=filtering.UrlFilter(lineAsArray3,delims,data)

    fmt.Printf("\nAfter UrlFilter:  delims: %v\n data: %v \n" , delims , data)

    // template_string:=filtering.GenTemplate(lineAsArray3, data)

    // fmt.Printf("###\nTemplate string:\n%s\n",template_string)

}

