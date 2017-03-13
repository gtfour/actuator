package analyze

var FOUND_IS_EMPTY         int = -4004

var DIGIT_LESS_INTERVAL    int = 3579
var DIGIT_GREATER_INTERVAL int = 9753
var DIGIT_IN_INTERVAL      int = 9779


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
