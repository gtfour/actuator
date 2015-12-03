package main
import "client_side/cuda"
import "fmt"
//import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {
    /*line:=`"UUID=6904489d-da97-46be-bee7-7ffdac09cf20" /home/user/Downloads/ ext3 errors=remount-ro 0       1`
    fmt.Printf("%s",line)
    lineAsArray:=strings.Split(line,"")
    quota_indexes:=cuda.GetQuotesIndexes(line)
    fmt.Printf("\n quota indexes:  %v\n",quota_indexes)
    pairs:=cuda.GroupByQuotes(lineAsArray, quota_indexes)
    for i:=range pairs {
         line_chunk:=line[pairs[i][0]+1:pairs[i][1]]
         fmt.Printf("\nSend to equalsign: %s\n",line_chunk)
         escape_eqal:=cuda.EqualSignEscape(line_chunk)
         fmt.Printf("\n escape_eqal:  %v\n",escape_eqal)
         for z:= range escape_eqal {
             fmt.Printf("\n%v\n", line_chunk[escape_eqal[z][0]:escape_eqal[z][1]])
         }
    }
    fmt.Printf("\npairs:  %v\n",pairs)*/
    line:="1:hello:::test:2"

    cuda.DebugPrintCharCounter(line)

    colon_indexes:=cuda.Escape_Colon(line)
    for i := range colon_indexes {
        start:=colon_indexes[i][0]
        end  :=colon_indexes[i][1]
        if start <= end {
            fmt.Printf("%s\n",line[start:end+1])
        } else {
            fmt.Printf("<None>\n")
        }
    }

    fmt.Printf("\n%v",colon_indexes)

}
