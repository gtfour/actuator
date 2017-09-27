package filtering

//import "fmt"
import "jumper/cuda/analyze"

func BaseFilter(lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int){
    //
    // Simple wrapper
    //
    return
}

func BracketsFilter(lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {
    //
    // var section_brackets_square   =  [2]string {"[","]"}
    // var section_brackets_triangle =  [3]string {"<",">","</"}
    // var section_brackets_curly    =  [2]string {"{","}"}
    //
    ndelims = delims
    ndata   = data
    return
    //
    //
    //
}

func SquareBracketsFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {
    square_brackets_open             := []string{"["}
    square_brackets_close            := []string{"]"}
    square_brackets_indexes          := analyze.ArrayInArrayIndexes(lineAsArray, square_brackets_open, square_brackets_close)
    square_brackets_complete_indexes := CombineDoubleSymbols(square_brackets_indexes)
    //fmt.Printf("\nSquare Brackets Complete Indexes: %v\n", square_brackets_complete_indexes)
    //ndelims = delims
    //ndata   = data
    //
    var data_inside_squauare_indexes [][]int
    if len(square_brackets_complete_indexes)>=1 {
        for sq:= range square_brackets_complete_indexes  {
            sq_range := square_brackets_complete_indexes[sq]
            if len(sq_range) == 2 {
                first_data                   := sq_range[0]+1
                last_data                    := sq_range[1]-1
                //
                // data between pair of square quotes
                //
                data_range                   := []int{first_data,last_data}
                data_inside_squauare_indexes =  append(data_inside_squauare_indexes, data_range)
                //
            }
        }
        ndelims,ndata = AlumaPaster(delims , data , Shifter(data_inside_squauare_indexes))
    } else {
        ndelims = delims
        ndata   = data
    }
    return
}


func DotFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {
    //
    //
    dot              := []string{"."}
    dot_indexes      := analyze.ArrayInArrayIndexes(lineAsArray, dot)
    new_data_indexes := analyze.GlueDataByConnector(data, dot_indexes)
    ndelims,ndata    =  AlumaPaster(delims , data , Shifter(new_data_indexes))
    //
    // debug
    //fmt.Printf("\n----\nGlueDataByConnector:\ndata: %v\ndot_indexes: %v\nnew_data: %v\n----\n",data,dot_indexes,analyze.GlueDataByConnector(data, dot_indexes))
    //
    return
}


func QuotesFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {

    single_quote := []string{"'"}
    double_quote := []string{`"`}
    grave_quote  := []string{"`"}


    single_quote_indexes := analyze.ArrayInArrayIndexes(lineAsArray, single_quote)
    double_quote_indexes := analyze.ArrayInArrayIndexes(lineAsArray, double_quote)
    grave_quote_indexes  := analyze.ArrayInArrayIndexes(lineAsArray, grave_quote)

    var quotes_complete_indexes   [][]int
    var data_inside_quote_indexes [][]int

    single_quote_complete_indexes := CombineDoubleSymbols(single_quote_indexes)
    double_quote_complete_indexes := CombineDoubleSymbols(double_quote_indexes)
    grave_quote_complete_indexes  := CombineDoubleSymbols(grave_quote_indexes)
    //
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
                    // wtf ?!
                    data_inside_quote_indexes=append(data_inside_quote_indexes, new_quote_range)
                } else if last_data < first_data  { // example: line3:=`cache_file_prefix = ""`
                    // wtf ?!
                    data_inside_quote_indexes=append(data_inside_quote_indexes, new_quote_range)
                }
            }
        }
        // fmt.Printf("\nstrada to Aluma %v\n",data_inside_quote_indexes)
        ndelims,ndata = AlumaPaster(delims , data , Shifter(data_inside_quote_indexes))
    } else {
        ndelims = delims
        ndata   = data
    }
    //fmt.Printf("s: %v d: %v g: %v", single_quote_complete_indexes, double_quote_complete_indexes, grave_quote_complete_indexes )
    return
}

func PathFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {
    //
    // PATH_SPEC_CHARS   := []string {"/"}
    //
    //
    PATH_SPEC_CHARS     := []string{"%",":","/","@","?","#","-",".","_","+","="}
    path_marker         := []string{"/"}
    path_marker_indexes := analyze.ArrayInArrayIndexes(lineAsArray,path_marker)
    //
    //
    //
    if len(path_marker_indexes) > 0 {
        var path_complete_indexes [][]int
        for i := range path_marker_indexes {
            //
            path_index := path_marker_indexes[i]
            //
            if len(path_index) !=2 { continue }
                leftSearcher          := Searcher{ direction:LEFT_DIRECTION }
                leftSearcher.since    =  path_index[0]
                leftSearcher.accepter =  func ( char string )( bool ) {
                                         return false
                                     }
                rightSearcher          := Searcher{ direction:RIGHT_DIRECTION }
                rightSearcher.since    =  path_index[1]
                rightSearcher.accepter =  func ( char string )( bool ){
                                          return analyze.IsUnicodeLetter( char ) || analyze.IsUnicodeDigit( char ) || analyze.IsSymbolIn( char, PATH_SPEC_CHARS )
                                     }
                //
                rightSearcher.breaker  = func ( char string )( bool ){
                                         if char == ":/" {  return true } else { return false }
                                     }
                rightSearcher.breakerInputSize = 2
                //
                searchers:=[]Searcher { rightSearcher, leftSearcher }
                new_indexes:=RunSearchers( lineAsArray, searchers )
                //
                path_complete_indexes = append( path_complete_indexes, new_indexes )

        }
        // fmt.Printf("\n:filtering:PathFilter:\npath_complete_indexes:Before:\n%v\n",path_complete_indexes)
        path_complete_indexes = Shifter( path_complete_indexes )
        // fmt.Printf("\n:filtering:PathFilter:\nShifter(path_complete_indexes):After:\n%v\n",path_complete_indexes)
        // fmt.Printf("\n---\nBefore Aluma: Delims:%v\tDatas:%v\n---\n",delims , data  )
        ndelims,ndata         = AlumaPaster( delims , data , path_complete_indexes )
        // fmt.Printf("\n---\nAfter Aluma: Delims:%v\tDatas:%v\n---\n", ndelims,ndata )
        //
        ndelims = Shifter(ndelims)
        ndata   = Shifter(ndata)
        //
    } else {
        ndelims = delims
        ndata   = data
    }
    //
    //fmt.Printf("\nPathFilter: lineAsArray %v\n delims %v\n data %v\nndelims %v\n ndata %v\n",lineAsArray,delims,data,ndelims,ndata)

    //
    return ndelims, ndata
    //
    //
    //
}

func UrlFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {


    URL_SPEC_CHARS     := []string{"%","=",":","/","@","?","#","-",".","_","$"} // $ for baseurl=http://vault.centos.org/7.0.1406/extras/$basearch/
    url_marker_short   := []string{ ":","/","/"     }
    url_marker_long    := []string{ ":","/","/","/" }
    //
    url_marker_indexes := analyze.ArrayInArrayIndexes(lineAsArray,url_marker_short,url_marker_long)
    //
    // seem's cause a bug because in this case we have duplicate indexes 
    // Hope  if we push both indexes inside single array Shifter may clean nested indexes
    // solution is below:
    //
    // fmt.Printf("\n:: UrlFilter::url_marker_indexes::\n%v ::\n", url_marker_indexes)
    // url_marker_indexes = Shifter( url_marker_indexes )
    // fmt.Printf("\n:: UrlFilter::url_marker_indexes::after Shifter()\n%v ::\n", url_marker_indexes)
    // url_marker_indexes = Shifter(url_marker_indexes)   :) don't impact to anything , still have a bug
    //
    //
    if len(url_marker_indexes)>0 {
        var url_complete_indexes [][]int
        for i := range url_marker_indexes {
            url_index:=url_marker_indexes[i]
            if len(url_index)!=2 { continue }
            leftSearcher          := Searcher{direction:LEFT_DIRECTION}
            leftSearcher.since    =  url_index[0]
            leftSearcher.accepter =  func (char string)(bool) {
                                         return analyze.IsUnicodeLetter(char) || analyze.IsUnicodeDigit(char)
                                     }

            rightSearcher          := Searcher{direction:RIGHT_DIRECTION}
            rightSearcher.since    =  url_index[1]
            rightSearcher.accepter = func (char string)(bool) {
                                         return analyze.IsUnicodeLetter(char) || analyze.IsUnicodeDigit(char) || analyze.IsSymbolIn(char, URL_SPEC_CHARS )
                                     }
            searchers             :=  []Searcher {rightSearcher, leftSearcher}
            new_indexes           :=  RunSearchers(lineAsArray, searchers)
            url_complete_indexes  =   append(url_complete_indexes, new_indexes)
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

    //fmt.Printf("\nUrlFilter: lineAsArray %v\n delims %v\n data %v\nndelims %v\n ndata %v\n",lineAsArray,delims,data,ndelims,ndata)

    return ndelims,ndata

}

