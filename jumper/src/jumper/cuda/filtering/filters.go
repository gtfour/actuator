package filtering

import "fmt"
import "jumper/cuda"

func BaseFilter(lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int){
    // Simple wrapper
    return
}

func BracketsFilter(lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {
    //var section_brackets_square   =  [2]string {"[","]"}
    //var section_brackets_triangle =  [3]string {"<",">","</"}
    //var section_brackets_curly    =  [2]string {"{","}"}
    return

}


func QuotesFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {

    single_quote :=[]string{"'"}
    double_quote :=[]string{`"`}
    grave_quote  :=[]string{"`"}


    single_quote_indexes := cuda.ArrayInArrayIndexes(lineAsArray, single_quote)
    double_quote_indexes := cuda.ArrayInArrayIndexes(lineAsArray, double_quote)
    grave_quote_indexes  := cuda.ArrayInArrayIndexes(lineAsArray, grave_quote)

    var quotes_complete_indexes   [][]int
    var data_inside_quote_indexes [][]int
    single_quote_complete_indexes := CombineDoubleSymbols(single_quote_indexes)
    double_quote_complete_indexes := CombineDoubleSymbols(double_quote_indexes)
    grave_quote_complete_indexes  := CombineDoubleSymbols(grave_quote_indexes)
    for s:= range single_quote_complete_indexes {
        single:=single_quote_complete_indexes[s]
        quotes_complete_indexes=append(quotes_complete_indexes, single)
    }
    //
    for d:= range double_quote_complete_indexes {
        double:=double_quote_complete_indexes[d]
        quotes_complete_indexes=append(quotes_complete_indexes, double)
    }
    //
    for g:= range grave_quote_complete_indexes {
        grave:=single_quote_complete_indexes[g]
        quotes_complete_indexes=append(quotes_complete_indexes, grave)
    }
    //
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

func PathFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {
    //PATH_SPEC_CHARS     :=[]string {"/"}
    PATH_SPEC_CHARS       :=[]string{"%",":","/","@","?","#","-",".","_","+","="}
    path_marker           :=[]string {"/"}
    path_marker_indexes   :=cuda.ArrayInArrayIndexes(lineAsArray,path_marker)

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
                                         return cuda.IsUnicodeLetter(char) || cuda.IsUnicodeDigit(char) || cuda.IsSymbolIn(char, PATH_SPEC_CHARS )
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


    URL_SPEC_CHARS     := []string{"%","=",":","/","@","?","#","-",".","_","$"} // $ for baseurl=http://vault.centos.org/7.0.1406/extras/$basearch/
    url_marker_short   := []string{":","/","/"}
    url_marker_long    := []string{":","/","/","/"}
    url_marker_indexes := cuda.ArrayInArrayIndexes(lineAsArray,url_marker_short,url_marker_long)

    if len(url_marker_indexes)>0 {
        var url_complete_indexes [][]int
        for i := range url_marker_indexes {
            url_index:=url_marker_indexes[i]
            if len(url_index)!=2 { continue }
            leftSearcher          := Searcher{direction:LEFT_DIRECTION}
            leftSearcher.since    =  url_index[0]
            leftSearcher.accepter =  func (char string)(bool) {
                                         return cuda.IsUnicodeLetter(char) || cuda.IsUnicodeDigit(char)
                                     }

            rightSearcher          := Searcher{direction:RIGHT_DIRECTION}
            rightSearcher.since    =  url_index[1]
            rightSearcher.accepter = func (char string)(bool) {
                                         return cuda.IsUnicodeLetter(char) || cuda.IsUnicodeDigit(char) || cuda.IsSymbolIn(char, URL_SPEC_CHARS )
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

