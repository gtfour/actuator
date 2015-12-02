package main
import "client_side/cuda"
import "fmt"
//import "strings"
// split each file and each line in file by space and save into slice of byte()  slices


func main() {
          
    /*line:=`"UUID=6904489d-da97-46be-bee7-7ffdac09cf20" /home/user/Downloads/ ext3 errors=remount-ro 0       1`
    fmt.Printf("%s",line)
    heads,foots:=cuda.DebugCharCounter(line)
    for i:=range heads {
        fmt.Printf("\n%s\n%s\n",heads[i],foots[i])
    }
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
    _,_,_=cuda.SectionNameEscape("DEFAULT]")
    fmt.Printf("%s",cuda.RemoveSpaces(`  <Directory "/THDL/thdl-site">       `,1))

}
