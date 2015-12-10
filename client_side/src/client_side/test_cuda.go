package main
import "client_side/cuda"
import "fmt"
import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {

    line:="  2  2 2 2   2 22    32 32 3 23 2 32                               333"
    //sign_map:=cuda.SignMap()
    cuda.DebugPrintCharCounter(line)
    fmt.Printf("\n === \n")
    lineAsArray:=strings.Split(line, "")
    fmt.Printf(cuda.RemoveDupSpaces(lineAsArray))

}
