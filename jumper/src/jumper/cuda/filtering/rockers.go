package filtering

import "fmt"
//import "strings"
import "jumper/cuda/analyze"

func EqualSignRocker(lineAsArray []string, delims [][]int, data [][]int)(dataKey [][]int, dataValue [][]int){

    //

    equal_sign         := []string{ "=" }
    equal_sign_indexes := analyze.ArrayInArrayIndexes(lineAsArray, equal_sign)

    //

    fmt.Printf("\n equal_sign_indexes: %v len(equal_sign_indexes): %v\n",equal_sign_indexes, len(equal_sign_indexes))
    return

    //

}
