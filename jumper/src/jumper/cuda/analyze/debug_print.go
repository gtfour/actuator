package analyze

import "fmt"
import "strings"

func DebugCharCounter (line  string) (heads, foots []string) {
    lineAsArray := strings.Split(line,"")
    head        := ""
    foot        := ""
    counter     := 0
    //for i:=0 ; i<len(lineAsArray) ; i++ {
    for _,c := range lineAsArray {
        delim:=""
        delim_template:=" %s%s "
        for z:=2;z<=len(fmt.Sprint(counter));z++ {delim+=" "}
        head+=fmt.Sprintf(delim_template, string(c), delim)
        foot+=fmt.Sprintf("|%d|",counter)
        if (counter%10==0)&&(counter!=0) || (counter+1==len(lineAsArray))  { heads=append(heads,head) ; foots=append(foots,foot) ; head="" ; foot="" }
        counter+=1
    }
    return heads, foots
}

func DebugPrintCharCounter (line string) {
    heads,foots:=DebugCharCounter(line)
    for i:=range heads {
        fmt.Printf("\n%s\n%s\n",heads[i],foots[i])
    }
}
