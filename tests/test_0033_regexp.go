package main
import "regexp"
import "fmt"

func main() {


    test:="{Hello}"
    nRx:= regexp.MustCompile("[{|}]")
    var result = [][]int {}
    result=nRx.FindAllStringIndex(test, -1)
    fmt.Printf("\n%v\n",result)


}
