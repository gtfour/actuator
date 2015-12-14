package main
import "client_side/cuda"
import "fmt"
import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {

    //line:="  2  2 2 2   2 22    32 32 3 23 2 32                               333"
    line2:="hello = 1"
    lineAsArray2:=strings.Split(line2, "")
    //sign_map:=cuda.SignMap()
    cuda.DebugPrintCharCounter(line2)
    fmt.Printf("\n === \n")
    //lineAsArray:=strings.Split(line, "")
    //fmt.Printf(cuda.RemoveDupSpaces(lineAsArray))
    parser:=cuda.MakeParser("=")

    fmt.Printf("%v",parser(lineAsArray2))
    fmt.Printf("\n --- \n")
    indexes:=cuda.Escape_EqualSign(lineAsArray2)
    first_part:=indexes[0]
    second_part:=indexes[1]
    fmt.Printf("\n --- \n")
    fmt.Printf("%s",line2[first_part[0]:first_part[1]])
    fmt.Printf("%s",line2[second_part[0]:second_part[1]])
    fmt.Printf("\n-----\n")
    fmt.Printf("\n --- \n %v",cuda.PrepareData(lineAsArray2))

}
