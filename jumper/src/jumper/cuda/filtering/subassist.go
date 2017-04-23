package filtering
//
//
//
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
