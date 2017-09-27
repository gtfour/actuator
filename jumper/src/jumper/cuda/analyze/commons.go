package analyze

var FOUND_IS_EMPTY         int = -4004

var DIGIT_LESS_INTERVAL    int = 3579
var DIGIT_GREATER_INTERVAL int = 9753
var DIGIT_IN_INTERVAL      int = 9779


func ArrayInArrayIndexes (abc []string, phrases ...[]string )(indexes [][]int) {
    // 
    // have a bug caused by this function 
    //
    //
    if ( len(abc) < 1 )||( len(phrases) < 1 ){ return }
    for i := 0; i < len(abc); i++  {
        symbol := abc[i]
        var found [][]int
        for p := range phrases {
            phrase      :=  phrases[p]
            local_found :=  make( []int, 2 )
            if len(phrase) > 1 {
                    zsymbol := phrase[0]
                    if symbol == zsymbol {
                        local_found[0] =  i
                        counter        := 1    // 1 because we are  keeping in mind that we already checked zsymbol with symbol earlier
                        xi             := i+1  // here is a bug : xi increased without checking len with abc : panic: runtime error: index out of range 
                                               // bug may cause by running over string:  # To use this repo, put in your DVD and use it with the other repos too:
                        for  ; counter < len( phrase ) ; {
                            if xi >= len(abc) { break } // fix
                            xsymbol    := abc[xi] // xi may be bigger than abc size this is "cause index out of range" error
                            if xsymbol != phrase[counter] { break }
                            if counter >= len(phrase)-1 { local_found[1] = xi ; found = append(found, local_found) ; break }
                            xi         += 1
                            counter    += 1
                        }
                    } else {
                        // ??? add continue or not
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

func DigitInInterval(digit int, interval []int)(int){
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

func GlueDataByConnector(data_indexes [][]int, connector_indexes [][]int)(new_data_indexes [][]int){
    //
    // seems will not with following string ".hello"
    //
    data_indexes_length := len(data_indexes)
    last_elem           := false
    skip_next           := false
    currentDataIndex    := make([]int,0)
    connector_x         := 0
    //
    //
    for i := range data_indexes {
        dataIndex           := data_indexes[i]
        // newDataIndex     := make([]int,0)
        // currentDataIndex := make([]int,0)
        // newDataIndex     =  dataIndex
        if !skip_next {
            currentDataIndex = dataIndex
        }
        if i == data_indexes_length-1 { last_elem=true }
        if !last_elem {
            dataIndexNext := data_indexes[i+1]
            //for x := range connector_indexes {
            pair_found := false
            for x := connector_x; x<len(connector_indexes); x++ {
                connectorIndex := connector_indexes[x]
                if len(connectorIndex) == 2 && len(dataIndex) == 2 && len(dataIndexNext) == 2 {
                    connectorIndexFirst   := connectorIndex[0]
                    connectorIndexLast    := connectorIndex[1]
                    //currentDataIndexFirst := currentDataIndex[0]
                    currentDataIndexLast  := currentDataIndex[1]
                    dataIndexNextFirst    := dataIndexNext[0]
                    dataIndexNextLast     := dataIndexNext[1]
                    if currentDataIndexLast == connectorIndexFirst-1 && dataIndexNextFirst == connectorIndexLast+1 {
                        currentDataIndex[1] = dataIndexNextLast
                        skip_next           = true
                        pair_found          = true
                        connector_x = x // start since this position next time
                        break
                    }
                }
            }
            if !pair_found {
                skip_next        = false
                new_data_indexes = append(new_data_indexes, currentDataIndex)
            }
        } else {
            new_data_indexes = append(new_data_indexes, currentDataIndex)
        }
    }
    //
    return
}
