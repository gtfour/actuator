package cuda
import "fmt"

var left_direction   int = 1100
var right_direction  int = 1001
var both_directions  int = 2002
var FOUND_IS_EMPTY   int = -4004

type Symbol struct {
    Value           string
    SearchDirection int
    MaxCount        int
    Accepter        func(string)(bool)
    Breaker         func(string)(bool)
}

var URL_SPEC_CHARS = []string {"%","=",":","/","@","?","#"}

type Cyclone struct {
    // line prop
}

func DataHeaderSelector(first_table_string []string)(data [][]int, isTableHeader bool ) {
    return
}

/*func UrlSelector(str []string, delim []int,  data_before []int , data_after []int)(data [][]int, isUrl bool ) {
    fmt.Printf("Delim:%v StrPart:%v", delim,str[delim[0]:delim[1]])
    return
}

func UrlMatcher(str []string, delim []int ) {

    match:=str[delim[0]:delim[1]]
    fmt.Printf("match:%v  str:%v  delim:%v",match,str,delim )


}*/

func StringArrayIsEqual (abc , def []string) (bool) {

    return true


}

func UrlFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {

    fmt.Printf("\n line: %v \n  delims: %v \n  data: %v \n",lineAsArray,delims,data)
    return ndelims,ndata

}

func ArrayInArrayIndexes (abc []string, phrases ...[]string )(indexes [][]int) {

    if (len(abc) < 1 )||(len(phrases) < 1){return}
    for i := 0; i < len(abc); i++  {
        symbol:=abc[i]
        var found [][]int
        for p := range phrases {
            local_found:=make( []int, 2 )
            phrase:=phrases[p]
            if len(phrase) > 1 {
                    zsymbol := phrase[0]
                    if symbol == zsymbol {
                        local_found[0] =  i
                        counter        := 1
                        xi             := i+1
                        for  ; counter < len(phrase) ;  {
                            xsymbol    := abc[xi]
                            if xsymbol != phrase[counter] { break }
                            if counter >= len(phrase)-1 { local_found[1] = xi ; found = append(found, local_found) ; break }
                            xi         += 1
                            counter    += 1
                        }
                    } else {
                    }
            }
        }
        arrayWithMaxLenIndex:=CompareArrayLen(found)
        if arrayWithMaxLenIndex != FOUND_IS_EMPTY {
            indexes = append(indexes, found[arrayWithMaxLenIndex])
        }
    }
    return
}

func CompareArrayLen (indexes [][]int)(int) {

    var max_len       int
    var max_len_index int
    for i := range indexes {
        array:=indexes[i]
        if len(array) ==2 {
            first:=array[0]
            last :=array[1]
            diff := last-first
            if diff >= max_len{
                max_len       = diff
                max_len_index = i
            }
        } else { continue }
    }
    if max_len == 0 { return FOUND_IS_EMPTY }
    return max_len_index
}

