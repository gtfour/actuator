package filtering

import "jumper/cuda/analyze"

func ColonFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int) {
    //
    // joining words splitted by colon 
    //
    colon          :=  []string{":"}
    // colonsCount    :=  len(colon)
    // _              =   colonsCount
    colon_indexes  :=  analyze.ArrayInArrayIndexes(lineAsArray, colon)
    //
    if len( colon_indexes ) >= 1 {
        data          := analyze.MakeReverse( colon_indexes, len(lineAsArray) )
        dataShifted   := Shifter(data)
        ndelims,ndata =  AlumaPaster(delims, data, dataShifted)
    } else {
        ndelims = delims
        ndata   = data
    }
    return
}

func EqualSignFilter( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int){
    //
    // joining words splitted by colon 
    //
    esign            :=  []string{"="}
    // colonsCount   :=  len(sign)
    // colon_indexes :=  analyze.ArrayInArrayIndexes(lineAsArray, signs)
    esign_indexes    :=  analyze.ArrayInArrayIndexes(lineAsArray, esign)
    //
    // split current line on two parts by equal sign
    //
    if len(esign_indexes) == 1 {
        data           := analyze.MakeReverse(esign_indexes, len(lineAsArray))
        dataShifted    := Shifter(data)
        ndelims,ndata  =  AlumaPaster(delims, data, dataShifted)
    } else {
        ndelims = delims
        ndata   = data
    }
    return
    //
    //
    //
}
