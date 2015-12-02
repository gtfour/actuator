package main
import "client_side/cuda"
import "fmt"
import "strings"


func main() {

    line:=`UUID="6904489d-da97-46be-bee7-7ffdac09cf20" /home/user/Downloads/ ext3 errors=remount-ro 0       1`
    lineAsArray:=strings.Split(line,"")
    quota_indexes:=cuda.GetQuotesIndexes(line)
    fmt.Printf("\n quota indexes:  %v\n",quota_indexes)
    pairs:=cuda.GroupByQuotes(lineAsArray, quota_indexes)
    fmt.Printf("\npairs:  %v\n",pairs)

}
