package cuda
import "fmt"

var LEFT_DIRECTION         int = 1100
var RIGHT_DIRECTION        int = 1001
var DIGIT_LESS_INTERVAL    int = 3579
var DIGIT_GREATER_INTERVAL int = 9753
var DIGIT_IN_INTERVAL      int = 9779

//var BOTH_DIRECTIONS  int = 2002
var FOUND_IS_EMPTY         int = -4004

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

func BracketsFilter(lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {

    var section_brackets_square   =  [2]string {"[","]"}
    var section_brackets_triangle =  [3]string {"<",">","</"}
    var section_brackets_curly    =  [2]string {"{","}"}

}


func QuotesFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {

    single_quote :=[]string{"'"}
    double_quote :=[]string{`"`}
    grave_quote  :=[]string{"`"}


    single_quote_indexes := ArrayInArrayIndexes(lineAsArray, single_quote)
    double_quote_indexes := ArrayInArrayIndexes(lineAsArray, double_quote)
    grave_quote_indexes  := ArrayInArrayIndexes(lineAsArray, grave_quote)

    var quotes_complete_indexes   [][]int
    var data_inside_quote_indexes [][]int
    single_quote_complete_indexes := CombineDoubleSymbols(single_quote_indexes)
    double_quote_complete_indexes := CombineDoubleSymbols(double_quote_indexes)
    grave_quote_complete_indexes  := CombineDoubleSymbols(grave_quote_indexes)
    for s:= range single_quote_complete_indexes {
        single:=single_quote_complete_indexes[s]
        quotes_complete_indexes=append(quotes_complete_indexes, single)
    }
    for d:= range double_quote_complete_indexes {
        double:=double_quote_complete_indexes[d]
        quotes_complete_indexes=append(quotes_complete_indexes, double)
    }
    for g:= range grave_quote_complete_indexes {
        grave:=single_quote_complete_indexes[g]
        quotes_complete_indexes=append(quotes_complete_indexes, grave)
    }
    if len(quotes_complete_indexes)>=1 {
        for q:= range quotes_complete_indexes {
            quote_range:=quotes_complete_indexes[q]
            if len(quote_range) == 2 {
                first_data := quote_range[0]+1
                last_data  := quote_range[1]-1
                new_quote_range:=[]int{first_data,last_data}
                if first_data < last_data {
                    data_inside_quote_indexes=append(data_inside_quote_indexes, new_quote_range)
                } else if last_data < first_data  { // example: line3:=`cache_file_prefix = ""`
                    data_inside_quote_indexes=append(data_inside_quote_indexes, new_quote_range)
                }
            }
        }
        fmt.Printf("\nstrada to Aluma %v\n",data_inside_quote_indexes)
        ndelims,ndata = AlumaPaster(delims , data , Shifter(data_inside_quote_indexes))
    } else {
        ndelims = delims
        ndata   = data
    }
    //fmt.Printf("s: %v d: %v g: %v", single_quote_complete_indexes, double_quote_complete_indexes, grave_quote_complete_indexes )
    return
}

func PathFilter ( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {
    //PATH_SPEC_CHARS     :=[]string {"/"}
    PATH_SPEC_CHARS     :=[]string{"%",":","/","@","?","#","-",".","_","+","="}
    path_marker         := []string {"/"}
    path_marker_indexes := ArrayInArrayIndexes(lineAsArray,path_marker)
    //fmt.Printf("PathFilter %v lineAsArray %v ", path_marker_indexes , lineAsArray)
    if len(path_marker_indexes)>0 {
        var path_complete_indexes [][]int
        for i := range path_marker_indexes {
            path_index         :=  path_marker_indexes[i]
            if len(path_index) !=2 { continue }
                leftSearcher          := Searcher{direction:LEFT_DIRECTION}
                leftSearcher.since    =  path_index[0]
                leftSearcher.accepter =  func (char string)(bool) {
                                         return false
                                     }
                rightSearcher          := Searcher{direction:RIGHT_DIRECTION}
                rightSearcher.since    =  path_index[1]
                rightSearcher.accepter = func (char string)(bool) {
                                         return IsUnicodeLetter(char) || IsUnicodeDigit(char) || IsSymbolIn(char, PATH_SPEC_CHARS )
                                     }
                searchers:=[]Searcher {rightSearcher, leftSearcher}
                new_indexes:=RunSearchers(lineAsArray, searchers)
                path_complete_indexes = append(path_complete_indexes, new_indexes)

        }
        path_complete_indexes = Shifter(path_complete_indexes)
        ndelims,ndata = AlumaPaster(delims , data , path_complete_indexes)
    } else {
        ndelims = delims
        ndata   = data
    }
    return ndelims,ndata
}

func UrlFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {


    URL_SPEC_CHARS     :=[]string{"%","=",":","/","@","?","#","-",".","_","$"} // $ for baseurl=http://vault.centos.org/7.0.1406/extras/$basearch/
    url_marker_short   :=[]string{":","/","/"}
    url_marker_long    :=[]string{":","/","/","/"}
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
        ndelims,ndata = AlumaPaster(delims , data , url_complete_indexes)
        // ex
        ndelims=Shifter(ndelims)
        ndata  = Shifter(ndata)
        //
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
            } else if len(phrase) == 1  {
                zsymbol := phrase[0]
                if symbol == zsymbol {
                    local_found[0] = i
                    local_found[1] = i
                    found = append(found, local_found)
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
    max_len := -1
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
    if max_len == -1 { return FOUND_IS_EMPTY }
    return max_len_index
}

func RunSearchers(lineAsArray []string,searchers []Searcher)( extended_indexes []int  ) {
    extended_indexes = make([]int, 2)
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

func AlumaPaster (delims [][]int, data [][]int, strada [][]int) (ndelims [][]int, ndata [][]int) {
    // strada should be inserted in data array
    // delims with indexes included in strada will be ignored
    // data  with indexes included  in strada will be ignored
    //fmt.Printf("  delims: %v\n  data: %v\n strada: %v\n",delims,data,strada)
    var last_delim_index int
    var last_data_index  int
    delims_last_elem := delims[(len(delims)-1)]
    data_last_elem   := data[(len(data)-1)]
    if len(delims_last_elem)==2 && len(data_last_elem)==2 {
        last_delim_index = delims_last_elem[1]
        last_data_index  = data_last_elem[1]
    }
    for i := range strada {
        ndelims := [][]int {}
        indexes:=strada[i]
        if len(indexes)!=2 { continue }  //{ break ; return delims, data }
        first := indexes[0]
        last  := indexes[1]
        //
        if first > last {
            first = indexes[1]
            last  = indexes[0]
        }
        //
        for de := range  delims {
            delim        := delims[de]
            if len(delim)!=2 { continue }
            first_delim  := delim[0]
            last_delim   := delim[1]
            first_state       := DigitInInterval(first, delim)
            last_state        := DigitInInterval(last, delim)
            first_delim_state := DigitInInterval(first_delim, indexes)
            last_delim_state  := DigitInInterval(last_delim,  indexes)
            //fmt.Printf("\nfirst %v | firststate %v | laststate %v | strada %v | delim %v | firstdelimstate %v | lastdelimstate %v \n ",first,first_state,last_state, strada,delim,first_delim_state, last_delim_state)
            if first_state == DIGIT_IN_INTERVAL && last_state == DIGIT_IN_INTERVAL {
               // split current delim to two new delims without strada indexes
               fmt.Printf("\nStrada on delim interval\n")
               new_delim_first := make([]int,2)
               new_delim_last  := make([]int,2)
               diff_first := first - first_delim
               diff_last  := last_delim  - last
               if diff_first>0 { new_delim_first[0] = first_delim ; new_delim_first[1] = first - 1 ; ndelims=append(ndelims, new_delim_first) }
               if diff_last >0 { new_delim_last[0]  = last +1     ; new_delim_last[1]  = last_delim; ndelims=append(ndelims, new_delim_last)  }
               //if diff_first == 0 && diff_last == 0 {   }
            } else if first_state == DIGIT_IN_INTERVAL {
                new_delim    := make([]int,2)
                diff_first   := first - first_delim
                if diff_first > 0{
                    new_delim[0]= first_delim
                    new_delim[1]= first-1
                    ndelims=append(ndelims, new_delim)
                }
            } else if last_state == DIGIT_IN_INTERVAL {
                new_delim    := make([]int,2)
                diff_last    := last_delim - last
                if diff_last > 0 {
                    new_delim[0]= last +1
                    new_delim[1]= last_delim
                    ndelims=append(ndelims, new_delim)
                }
            // ??? Pay attention
            } else if first_delim_state == DIGIT_IN_INTERVAL && last_delim_state == DIGIT_IN_INTERVAL {

            } else {
                ndelims=append(ndelims, delim)
            }
        }
        delims = ndelims
    }
    ndelims = delims

    last_matched_strada_id:=-1
    for da := range  data {
        data_part:=data[da]
        if len(data_part)!=2 { continue }
        first_data := data_part[0]
        last_data  := data_part[1]
        var includes      bool
        var insert_strada bool
        for i := range strada {
            indexes:=strada[i]
            if len(indexes)!=2 { continue }
            first_state          := DigitInInterval(first_data, indexes)
            last_state           := DigitInInterval(last_data, indexes)
            if first_state == DIGIT_IN_INTERVAL && last_state == DIGIT_IN_INTERVAL{
                includes = true
                if i != last_matched_strada_id {
                    insert_strada          = true
                    last_matched_strada_id = i
                }
            } else {
                fmt.Printf("\n::else::\n")
                interval_between_data:=make([]int,2)
                first_strada         := indexes[0]
                last_strada          := indexes[1]
                var replace_on_insert bool
                if first_strada > last_strada {
                    first_strada = indexes[1]
                    last_strada  = indexes[0]
                    replace_on_insert = true
                }
                //last_delim_index := delims[(len(delims)-1)][1]
                //last_data_index  := data[(len(data)-1)][1]
                if da == 0 {
                    interval_between_data[0] = 0
                    if len(data) == 1 {
                        if last_delim_index >= last_data_index {
                            interval_between_data[1] = last_delim_index
                        } else {
                            interval_between_data[1] = last_data_index
                        }
                    } else {
                        interval_between_data[1] = last_data

                    }
                } else if da == len(data)-1 {

                    interval_between_data[0] = first_data
                    if last_delim_index >= last_data_index {
                        interval_between_data[1]  = last_delim_index
                    } else {
                         interval_between_data[1] = last_data_index
                    }

                } else {
                    interval_between_data[0] = data_part[1]
                    interval_between_data[1] = data[da+1][0]
                }
                if  DigitInInterval(first_strada, interval_between_data)  == DIGIT_IN_INTERVAL && DigitInInterval(last_strada, interval_between_data)  == DIGIT_IN_INTERVAL {
                    last_matched_strada_id = i
                    if replace_on_insert == false {
                        insert_strada = true
                    } else {
                         nindexes:=make([]int,2)
                         nindexes[0] = last_strada
                         nindexes[1] = first_strada
                         insert_strada = true
                    }
                }
                //var interval_between_data []int
                //interval_between_data[0] 


            }
        }
        if includes == false {
            ndata=append(ndata, data_part)
        }
        if insert_strada == true {
            ndata=append(ndata, strada[last_matched_strada_id])
        }
    }
    return //ndelims,ndata
}

func DigitInInterval(digit int, interval []int) (int) {
    if digit <= interval[1] && digit >= interval[0] {
        return DIGIT_IN_INTERVAL
    }
    if digit < interval[0] {
        return DIGIT_LESS_INTERVAL
    }
    if digit > interval[1] {
        return DIGIT_GREATER_INTERVAL
    }
    return 0
}


func Shifter(interval [][]int)(ninterval [][]int) {
    var skipped []int
    for i:= range interval {
        if IsDigitIn(i,skipped) == false {
            parent_int_part:=interval[i]
            for z:= range interval {
                if z == i {continue} //do not compare interval with self
                int_part:=interval[z]
                if len(int_part)!=2{continue}

                first:=int_part[0]
                last :=int_part[1]

                if DigitInInterval(first,parent_int_part) == DIGIT_IN_INTERVAL && DigitInInterval(last,parent_int_part) == DIGIT_IN_INTERVAL {
                    skipped=append(skipped, z)
                }

            }

        }
    }
    for x:= range interval {
        if IsDigitIn(x,skipped) == false {
            ninterval=append(ninterval, interval[x])
        }
    }
    return
}

func CombineDoubleSymbols ( quotes_indexes [][]int  )(combined_quotes [][]int) {

    if len(quotes_indexes) >=2 && len(quotes_indexes)%2 == 0 {
        pair:=make([]int,2)
        pair[0] = -1
        pair[1] = -1
        for i:=0 ; i<len(quotes_indexes) ; i++ {
            index:=quotes_indexes[i]
            if pair[0] == -1 && len(index) == 2 && index[0] == index[1] {
                pair[0] = index[0]
            } else if pair[0] != -1 && pair[1] == -1 &&  len(index) == 2 && index[0] == index[1] {
                pair[1] = index[0]
                combined_quotes = append(combined_quotes, pair)
                pair = make([]int,2)
                pair[0] = -1
                pair[1] = -1
            }
        }
    } else {
        return combined_quotes
    }
    return
}
