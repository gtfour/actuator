package parse

func GetIndexes ( lineAsArray []string ) (delims [][]int , data [][]int) {
    var delimPair = []int {-1,-1}
    var dataPair  = []int {-1,-1}
    for i:= range lineAsArray {
        char:=lineAsArray[i]
        if IsSymbolIn(char,WORD_DELIM) == false && IsUnicodeLetter(char) == false  && IsUnicodeDigit(char) == false {
            if dataPair[0]  != -1 {
                dataPair[1] = i - 1 // make pair with previous element as second member of pair
                data      = append(data, dataPair)
                dataPair  = []int   {-1,-1}
            }
            if delimPair[0] == -1 {
                delimPair[0]= i
            }
            delimPair[1] = i
            if (i==(len(lineAsArray)-1)) || ((i<=len(lineAsArray)-2) && ( IsSymbolIn(lineAsArray[i+1],WORD_DELIM) == true || IsUnicodeLetter(lineAsArray[i+1]) == true  || IsUnicodeDigit(lineAsArray[i+1]) == true)) {
                    delims=append(delims, delimPair)
                    delimPair=[]int{-1, -1}
            }
        } else {
            if dataPair[0]  == -1 {
                dataPair[0] = i
            }
            if (i==(len(lineAsArray)-1)) {
                dataPair[1] = i
                data      = append(data, dataPair)
            }
        }
    }
    return delims,data
}
