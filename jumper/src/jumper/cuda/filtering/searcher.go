package filtering

type Searcher struct {
    // //
    // //
    value               string
    since               int
    direction           int
    maxCount            int
    // //
    // //  should satisfy to Accepter and Breaker . if  Accepter returns true and Breaker return false searching will remain
    accepter            func(string)(bool)
    breaker             func(string)(bool)
    resultPosition      int // calculating field
    // //
    // //
}


func RunSearchers(lineAsArray []string,searchers []Searcher)( extended_indexes []int  ) {
    extended_indexes = make([]int, 2)
    for sindex := range searchers {
        searcher := searchers[sindex]
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
