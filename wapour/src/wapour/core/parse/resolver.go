package parse

func GetIndexes ( lineAsArray []string ) (delims [][]int , data [][]int) {

    var delimPair = []int   {-1,-1}
    var dataPair  = []int   {-1,-1}
    //var offset int
    for i:= range lineAsArray {
        //offset = i
        char:=lineAsArray[i]
        //if IsSymbolIn(char,ABC,NUMBERS,WORD_DELIM) == false {
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
            //if ((i==(len(lineAsArray)-1)) || ((i<=len(lineAsArray)-2) && (IsSymbolIn(lineAsArray[i+1],ABC,NUMBERS,WORD_DELIM) == true)) ) {
            if (i==(len(lineAsArray)-1)) || ((i<=len(lineAsArray)-2) && ( IsSymbolIn(lineAsArray[i+1],WORD_DELIM) == true || IsUnicodeLetter(lineAsArray[i+1]) == true  || IsUnicodeDigit(lineAsArray[i+1]) == true)) {

                    // +1 because see /actuator/tests/test_0038_arr.go
                    // delimAsArray:=GetFixedArrayChars(lineAsArray[delimPair[0]:offset+1], []int { 0, (delimPair[1]-delimPair[0]) }) // have to add +1 .but   why !??!?
                    //delim_split_space:=Escape_Spaces(delimAsArray)
                    delims=append(delims, delimPair)
                    delimPair=[]int{-1, -1}
                    // simplifyig
                    /*
                    if len(delim_split_space) == 1 {
                        delims=append(delims, delimPair)
                        delimPair=[]int{-1, -1}
                    } else if len(delim_split_space) == 0   {
                       // it  seems that there are just a lot of spaces in delimAsArray  and nothing else
                        // delimPair[1] = delimPair[0] i am going change it to collect all space indexes instead of first
                        delims=append(delims, delimPair)
                        delimPair=[]int{-1, -1}
                    } else {
                        for sd := range delim_split_space {
                           delim_ss_pair:=delim_split_space[sd]
                           if len(delim_ss_pair) == 2 { delim_ss_pair[0]=delim_ss_pair[0]+delimPair[0] ; delim_ss_pair[1]=delim_ss_pair[1]+delimPair[0]  }
                           delims = append(delims, delim_ss_pair)
                        }
                        delimPair=[]int{-1, -1}
                    }*/

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

