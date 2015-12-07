package main
import "regexp"
import "fmt"

func main() {


    var test = []string { "[Hello]", "<Hello>" , "server { " }
    
    nRx:= regexp.MustCompile("[\\[|\\]|\\}|\\{|\\:|\\;|\\<|\\>|\\</]")
    var result = [][]int {}
    for i := range test {
        result=nRx.FindAllStringIndex(test[i], -1)
        fmt.Printf("\n%v\n", test[i])
        fmt.Printf("\n%v\n=====", result)
    }
    fmt.Printf("\n%v\n",result)


}
