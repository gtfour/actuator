package main
import "client_side/cuda"
import "fmt"
//import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {

    line:=""
    //sign_map:=cuda.SignMap()
    cuda.DebugPrintCharCounter(line)
    fmt.Printf("\n === \n")
    fmt.Printf(cuda.RemoveDupSpaces(line))


}
