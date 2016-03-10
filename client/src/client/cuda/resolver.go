package cuda
import "fmt"

var URL_SPEC_CHARS = []string {"%","=",":","/","@","?","#"}

type Cyclone struct {
    // line prop
}

func DataHeaderSelector(first_table_string []string)(data [][]int, isTableHeader bool ) {
    return
}

func UrlSelector(str []string, delim []int,  data_before []int , data_after []int)(data [][]int, isUrl bool ) {
    fmt.Printf("Delim:%v StrPart:%v", delim,str[delim[0]:delim[1]])
    return
}

func UrlMatcher(str []string, delim []int ) {

    match:=str[delim[0]:delim[1]]
    fmt.Printf("match:%v  str:%v  delim:%v",match,str,delim )


}

func StringArrayIsEqual (abc , def []string) (bool) {

    return true


}

func UrlFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {




    return ndelims,ndata

}





//func AnalyzeDelims()



