package main
import "client_side/cuda"
import "fmt"
import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {

    //line:="  2  2 2 2   2 22    32 32 3 23 2 32                               333"
    line2:="           name   =   [my_name]      "
    lineAsArray2:=strings.Split(line2, "")
    cuda.DebugPrintCharCounter(line2)
    parser:=cuda.MakeParser("[")

    //delims_indexes:=cuda.GetDelimsIndexes(lineAsArray2)

    
    fmt.Printf("\n parsed_string : %v \n",parser(lineAsArray2))

}
