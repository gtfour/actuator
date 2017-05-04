package filtering

import "fmt"
import "jumper/cuda/analyze"

func ColonFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {
    //
    // joining words splitted by colon 
    //
    colon          :=  []string{":"}
    colonsCount    :=  len(colon)
    _              =   colonsCount
    colon_indexes  :=  analyze.ArrayInArrayIndexes(lineAsArray, colon)
    //
    fmt.Printf("\nColon indexes: %v \n", colon_indexes)
    //
    if len( colon_indexes ) >= 1 {
        for c:= range quotes_complete_indexes {
            quote_range:=quotes_complete_indexes[q]
            if len(quote_range) == 2 {
                first_data := quote_range[0]+1
                last_data  := quote_range[1]-1
                new_quote_range:=[]int{first_data,last_data}
                if first_data < last_data {
                    data_inside_quote_indexes = append(data_inside_quote_indexes, new_quote_range)
                } else if last_data < first_data  { // example: line3:=`cache_file_prefix = ""`
                    data_inside_quote_indexes = append(data_inside_quote_indexes, new_quote_range)
                }
            }
        }
        // fmt.Printf("\nstrada to Aluma %v\n",data_inside_quote_indexes)
        ndelims,ndata = AlumaPaster(delims , data , Shifter(data_inside_quote_indexes))
    } else {
        ndelims = delims
        ndata   = data
    }
    //
    //
    //
    return
}
