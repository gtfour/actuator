package main
import "client_side/cuda"
import "fmt"
import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {

    line:=" NOZEROCONF=yes "
    //sign_map:=cuda.SignMap()
    indexes:=cuda.GetSignIndex(line)
    cuda.DebugPrintCharCounter(line)
    fmt.Printf("%v",indexes)
    if test, ok  :=indexes[cuda.EQUAL]; ok == true {
        fmt.Printf("\n====\n%v",cuda.GetSignScope(strings.Split(line,""), cuda.EQUAL , test[0] ))
    }


}
