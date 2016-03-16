package cuda
import "fmt"

var LEFT_DIRECTION   int = 1100
var RIGHT_DIRECTION  int = 1001
//var BOTH_DIRECTIONS  int = 2002
var FOUND_IS_EMPTY   int = -4004
var URL_SPEC_CHARS   = []string {"%","=",":","/","@","?","#","-",".","_"}

type SpecialWord struct {
    stype int // could be an email or url address or ip or  path
    pos   []int
    value string
}

type Searcher struct {
    value               string
    since               int
    direction           int
    maxCount            int
    // should satisfy to Accepter and Breaker . if  Accepter returns true and Breaker return false searching will remain
    accepter            func(string)(bool)
    breaker             func(string)(bool)
    resultPosition      int // calculating field
}


type LineAnalyzer struct {
    searchers     []Searcher
    specials      []SpecialWord
}


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

    url_marker_short :=[]string{":","/","/"}
    url_marker_long  :=[]string{":","/","/","/"}
    url_marker_indexes:=ArrayInArrayIndexes(lineAsArray,url_marker_short,url_marker_long)
    //fmt.Printf("\nurl_marker_indexes: %v\n", url_marker_indexes)
    if len(url_marker_indexes)>0 {
        var url_complete_indexes [][]int
        for i := range url_marker_indexes {
            url_index:=url_marker_indexes[i]
            if len(url_index)!=2 { continue }
            leftSearcher          := Searcher{direction:LEFT_DIRECTION}
            leftSearcher.since    =  url_index[0]
            leftSearcher.accepter =  func (char string)(bool) {
                                         return IsUnicodeLetter(char) || IsUnicodeDigit(char)
                                     }

            rightSearcher          := Searcher{direction:RIGHT_DIRECTION}
            rightSearcher.since    =  url_index[1]
            rightSearcher.accepter = func (char string)(bool) {
                                         return IsUnicodeLetter(char) || IsUnicodeDigit(char) || IsSymbolIn(char, URL_SPEC_CHARS )
                                     }
            searchers:=[]Searcher {rightSearcher, leftSearcher}
            new_indexes:=RunSearchers(lineAsArray, searchers)
            url_complete_indexes = append(url_complete_indexes, new_indexes)
        }
    } else {
        ndelims = delims
        ndata   = data
    }

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
    // return []int array index with max lenght 
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

func RunSearchers(lineAsArray []string,searchers []Searcher)( extended_indexes [2]int  ) {
    for sindex := range searchers {
        searcher:=searchers[sindex]
        if searcher.direction == RIGHT_DIRECTION && searcher.accepter!=nil {
            for i := searcher.since+1 ; i < len(lineAsArray); i++  {
                char:=lineAsArray[i]
                if searcher.accepter(char) == false {
                    extended_indexes[1] = i-1
                    break
                }
                if i == len(lineAsArray)-1 {
                    extended_indexes[1] = i
                    break
                }
            }
        } else if searcher.direction == LEFT_DIRECTION && searcher.accepter!=nil  {
            for i := searcher.since-1 ; i >=0  ; i--  {
                char:=lineAsArray[i]
                if searcher.accepter(char) == false {
                    extended_indexes[0] = i+1
                    break
                }
            }
        }
    }
    return
}

func AlumaPaster (delims [][]int, data [][]int, strada [][]int) (delims [][]int, data [][]int) {
    for i := range strada {
        indexes:=strada[i]
        if len(indexes)!=2 { continue }
        first := indexes[0]
        last  := indexes[1]
        //for  


    }
}
