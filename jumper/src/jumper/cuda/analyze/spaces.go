package analyze

import "strings"

func Escape_Spaces (oldlineAsArray []string) (indexes [][]int) {
    // 
    // Duplicate spaces will be present as one : wrong specification !!!
    //
    lineAsArray      := ReplaceTabsToSpaces(oldlineAsArray)
    var pair         =  []int { -1, -1 }
    for i := range lineAsArray {
        char:=lineAsArray[i]
        if char == " " && pair[0]!=-1 {
            pair[1] = i-1
            indexes = append( indexes, []int {pair[0], pair[1]})
            pair = []int { -1, -1 }
        } else if pair[0] == -1 && char != " " {
            pair[0] = i
        }
        if i == len(lineAsArray)-1 && pair[0]!=-1 {
            indexes = append( indexes, []int {pair[0], i})
        }
    }
    return indexes
}

func ReplaceTabsToSpaces ( lineAsArray []string ) ( newlineAsArray []string ) {

    entry:=strings.Join(lineAsArray, "")
    new_entry:=strings.Replace(entry, "	", " ", -1)
    return strings.Split(new_entry,"")

}
