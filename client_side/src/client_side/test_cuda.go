package main
import "client_side/cuda"
import "fmt"
import "strings"


func main() {

    line:=`"UUID=6904489d-da97-46be-bee7-7ffdac09cf20" /home/user/Downloads/ ext3 errors=remount-ro 0       1`
    fmt.Printf("\n Line:  %s\n",line)
    lineAsArray:=strings.Split(line,"")
    quota_indexes:=cuda.GetQuotesIndexes(line)
    fmt.Printf("\n quota indexes:  %v\n",quota_indexes)
    pairs:=cuda.GroupByQuotes(lineAsArray, quota_indexes)
    for i:=range pairs {
         fmt.Printf("\nSend to equalsign: %s\n",line[pairs[i][0]:pairs[i][1]])
         escape_eqal:=cuda.EqualSignEscape(line[pairs[i][0]:pairs[i][1]])
         fmt.Printf("\n escape_eqal:  %v\n",escape_eqal)
         for z:= range escape_eqal {
             fmt.Printf("\n%v\n", line[escape_eqal[z][0]:escape_eqal[z][1]])
         }
    }
    fmt.Printf("\npairs:  %v\n",pairs)

}
