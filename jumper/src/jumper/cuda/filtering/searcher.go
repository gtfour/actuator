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
    // accepterInputSize   int
    //
    breaker             func(string)(bool)
    breakerInputSize    int
    // //
    resultPosition      int // calculating field
    // //
    // //
}


func RunSearchers(lineAsArray []string,searchers []Searcher)( extended_indexes []int  ) {
    extended_indexes = make([]int, 2)
    for sindex := range searchers {
        searcher := searchers[sindex]
        if searcher.direction == RIGHT_DIRECTION && searcher.accepter != nil {
            //
            // will push additional condition check here , for correctly hanling of /etc/passwd
            //
            for i := searcher.since+1 ; i < len(lineAsArray); i++ {
                char := lineAsArray[i]
                //
                // checking breaker
                if searcher.breaker != nil {
                    // searcher.breaker(char)
                    //breakerInputSize := searcher.breakerInputSize
                }
                //
                //
                if searcher.accepter(char) == false {
                    extended_indexes[1] = i-1
                    break
                }
                if i == len(lineAsArray)-1 {
                    extended_indexes[1] = i
                    break
                }
            }
            //
            //
            //
        } else if searcher.direction == LEFT_DIRECTION && searcher.accepter != nil {
            //
            //
            //
            for i := searcher.since-1 ; i >=0  ; i--  {
                char:=lineAsArray[i]
                if searcher.accepter(char) == false {
                    extended_indexes[0] = i+1
                    break
                }
            }
            //
            //
            //
        } else if searcher.direction == RIGHT_DIRECTION && searcher.breaker!=nil {

        } else if searcher.direction == LEFT_DIRECTION && searcher.breaker!=nil {

        }
    }
    return
}

func PrepareCheckSet( lineAsArray []string, offset int, inputSize int )(checkSet string, err error){
    //
    //
    //
    if offset >= len(lineAsArray) || offset < 0 { return "",offset_out_of_range }
    // for i    := offset ; i < len(lineAsArray); i++ {
    lastChar := offset + inputSize - 1
    if lastChar < len(lineAsArray) {
        for i := offset ; i <= lastChar  ; i++ {
            char     := lineAsArray[i]
            checkSet =  checkSet+char
        }
    } else {
        return "", input_size_out_of_range
    }
    //}
    //
    //
    //
    return checkSet, nil
}
