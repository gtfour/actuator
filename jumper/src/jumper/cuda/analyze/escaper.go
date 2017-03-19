package analyze

import "strings"

func EscapeDoubleSign_functionBuilder( sign_start, sign_end string   ) ( doublesign_escaper  func ( lineAsArray []string ) ([][]int ) ) {

    // returns function to find single pair of text between specialized sign
    doublesign_escaper = func(lineAsArray []string) ( indexes [][]int) {
        pair := []int {-1,-1}
        for i:= range lineAsArray {
            char:=lineAsArray[i]
            if char == sign_start && pair[0] == -1  {
                pair[0] = i+1
                continue
            }
            if char == sign_end && pair[1] == -1  {
                pair[1] = i-1
                continue
            }
        }
        indexes=append(indexes, pair)
        return indexes
    }
    return doublesign_escaper
}

func EscapeSingleSign_functionBuilder( sign  string   ) ( singlesign_escaper  func ( lineAsArray []string ) ([][]int ) ) {

    // returns function to find single pair of text between specialized sign
    singlesign_escaper = func(lineAsArray []string) ( indexes [][]int) {

    entry:=strings.Join(lineAsArray, "")
    wordSplittedByEqualSign:= strings.Split(entry, sign)
    offset:=0
    for i := range wordSplittedByEqualSign {
        word:=wordSplittedByEqualSign[i]
        word_index:=strings.Index(entry[offset:], word)
        var new_pair =  []int {(word_index+offset),(len(word)-1+offset)}
        indexes=append(indexes, new_pair)
        offset+=len(word)+1 // 1 is equal sign

        }

        return indexes
    }
    return singlesign_escaper
}



func Escape_EqualSign (lineAsArray []string) (words_indexes [][]int) {

    entry:=strings.Join(lineAsArray, "")
    wordSplittedByEqualSign:= strings.Split(entry, "=")
    offset:=0
    for i := range wordSplittedByEqualSign {
        word:=wordSplittedByEqualSign[i]
        word_index:=strings.Index(entry[offset:], word)
        var new_pair =  []int {(word_index+offset),(len(word)-1+offset)}
        words_indexes=append(words_indexes, new_pair)
        offset+=len(word)+1 // 1 is equal sign
    }
    return words_indexes

}
