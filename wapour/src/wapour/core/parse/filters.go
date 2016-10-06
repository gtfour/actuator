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
    return
}
