package main
import "client/cuda"
//import "fmt"
import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {

    //line:="  2  2 2 2   2 22    32 32 3 23 2 32                               333"
    //line2:=`Default settings        secure_path="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"`
    //line3:= `    1353032691c5        ubuntuюペä:12.04        "/bin/bashробот"         7 weeks ago         Up 7 weeksペä                             prickly_leakeyペä    `
    //line3:="имя: Ваня item: Vodka"
    //line3:= `1353032691c5        ubuntu:12.04        "/bin/bash"         7 weeks ago         Up 7 weeks                              prickly_leakey`
    line3:=`deb http://repo.mongodb.org/apt/ubuntu trusty/mongodb-org/3.0 multiverse`
   // line3:="a: 122"
    //line3:=`a: 122`
    //line3:="                                     "
    //line3:="}}  }}"
    //lineAsArray2:=strings.Split(line2, "")
    //// lineAsArray3:=strings.Split(line3, "")
    lineAsArray3:=strings.Split(line3, "")
    cuda.DebugPrintCharCounter(line3)
    //parser:=cuda.MakeParser("[")

    delims,data:=cuda.GetIndexes(lineAsArray3)
    //fmt.Printf("\ndelims: %v\n data: %v \n" , delims , data)
    cuda.UrlFilter(lineAsArray3,delims,data)
    //cuda.UrlMatcher([]string {":","/","/"}, []int{0,1,2})

    /*for i := range delims_indexes{

        fmt.Printf("--\n%v\n--",cuda.GetFixedArrayChars(lineAsArray2, delims_indexes[i]))



    }*/
    //fmt.Printf("\nEscapeSpaces|%v| len|%d| \n",cuda.Escape_Spaces(lineAsArray3), len(cuda.Escape_Spaces(lineAsArray3)))

    /*data:=cuda.PrepareData(lineAsArray2, delims_indexes)

    fmt.Printf("\n data :\n %v \n",data)
    */
    //fmt.Printf("--%v--",cuda.Escape_Spaces(lineAsArray2))
}
