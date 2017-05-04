package analyze

func MakeReverse(inputIndexes [][]int, arraySize int)( outputIndexes [][]int) {
    //
    appropriateInterval  := []int{ 0,  arraySize-1 }
    //
    //
    reverseLeftPosition  := appropriateInterval[0]
    reverseRightPosition := 0
    lastInputIndex       := len(inputIndexes)-1
    lastArrayElemIndex   := arraySize-1
    //
    for i:= range inputIndexes {
        //
        outputPair                  := []int{ -1, -1 }
        inputPair                   := inputIndexes[i]
        if len(inputPair) != 2 { continue }
        first                       := inputPair[0]
        second                      := inputPair[1]
        reverseRightPosition        =  first - 1
        if reverseLeftPosition >= 0 && reverseRightPosition >=0 {
            outputPair[0] = reverseLeftPosition
            outputPair[1] = reverseRightPosition
            outputIndexes = append( outputIndexes, outputPair )
        }
        reverseLeftPosition = second + 1
        if i == lastInputIndex && second != lastArrayElemIndex {
            closePair     := []int {reverseLeftPosition, lastArrayElemIndex }
            outputIndexes =  append( outputIndexes, closePair)
        }
        //
    }
    return
    //
}
