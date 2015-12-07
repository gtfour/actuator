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


   /* line1:="		if a>2 && b==3{"
    line2:="if	a>2&&b==3{"
    line3:="	if a>2&&b==3	{	"
    line4:="	if a>2	&&	b==3	{"*/



    /*name:= []int {}
    tag:= []int {}
    var sectype int
    cuda.DebugPrintCharCounter(line1) ; name,tag,sectype=cuda.Escape_Section(line1)
    fmt.Printf("--\n%v\n%v\n%v\n--",name,tag,sectype)*/
    /*cuda.DebugPrintCharCounter(line2) ; name,tag,sectype=cuda.Escape_Section(line2)
    fmt.Printf("--\n%v\n%v\n%v\n--",name,tag,sectype)
    cuda.DebugPrintCharCounter(line3) ; name,tag,sectype=cuda.Escape_Section(line3)
    fmt.Printf("--\n%v\n%v\n%v\n--",name,tag,sectype)
    cuda.DebugPrintCharCounter(line4) ; name,tag,sectype=cuda.Escape_Section(line4)
    fmt.Printf("--\n%v\n%v\n%v\n--",name,tag,sectype)*/

    /*indexes:=cuda.Escape_Spaces(line1)
    fmt.Printf("\n")
    cuda.DebugPrintCharCounter(line1);fmt.Printf("%v",indexes)
    indexes=cuda.Escape_Spaces(line2)
    cuda.DebugPrintCharCounter(line2);fmt.Printf("%v",indexes)
    indexes=cuda.Escape_Spaces(line3)
    cuda.DebugPrintCharCounter(line3);fmt.Printf("%v",indexes)
    indexes=cuda.Escape_Spaces(line4)
    cuda.DebugPrintCharCounter(line4);fmt.Printf("%v",indexes)*/

    //indexes:=cuda.RemoveSpaces(line,2)
    //fmt.Printf("\n%v",indexes)
    //fmt.Printf("\n%s",line[indexes[0]:indexes[1]+1])
    /*line1:="hosts:          files mdns4_minimal [NOTFOUND=return] dns mdns4"
    cuda.DebugPrintCharCounter(cuda.ReplaceTabsToSpaces(line1))
    indexes:=cuda.Escape_Spaces(line1);fmt.Printf("%v",indexes)*/

    fmt.Printf("\n%v\n", cuda.SignMap())

}
