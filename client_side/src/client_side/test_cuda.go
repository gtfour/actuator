package main
import "client_side/cuda"
import "fmt"
//import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {
    line:="hello, hello 2/2*9 <Day>"
    indexes:=cuda.GetSignIndex(line)
    cuda.DebugPrintCharCounter(line)

    fmt.Printf("\n === \n")
    fmt.Printf("\n%v\n", indexes)

}
