package main
import "client_side/cuda"
import "fmt"
import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {

    //line:="  2  2 2 2   2 22    32 32 3 23 2 32                               333"
    //line2:="           name   =   [my_name]      "
    line3:="}}}}"
    lineAsArray3:=strings.Split(line3, "")
    cuda.DebugPrintCharCounter(line3)
    //parser:=cuda.MakeParser("[")

/*    delims_indexes:=cuda.GetDelimsIndexes(lineAsArray2)
    fmt.Printf("lineAsArray2:\n %v \ndelims_indexes:\n %v\n", lineAsArray2, delims_indexes)
    data:=cuda.PrepareData(lineAsArray2, delims_indexes)

    
    fmt.Printf("\n data :\n %v \n",data)
*/
    fmt.Printf("--%v--",cuda.Escape_Spaces(lineAsArray3))
}
