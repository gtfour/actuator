package handling

import "fmt"
import "jumper/cuda/analyze"

func GetSectionBreaker( line string, name []int, tag []int, sectionType int )(func(string)(bool)){
    //
    switch {

        case sectionType == analyze.SQUARE_SECTION:
            breaker := func(line string)(bool){
                //
                // this kind of section is closing by empty line
                //
                if line == "" {
                    return true
                } else {
                    return false
                }
            }
            return breaker
        case sectionType == analyze.TRIANGLE_SECTION_STARTING:
            breaker := func(phrase string)(bool){
                if len(name) == 2 {
                    match := "</" + line[name[0]:name[1]+1]+">"
                    fmt.Printf("\n --- triangle match: %v --- \n",match)
                    if phrase == match { return true } else { return false }
                } else {
                    return false
                }
            }
            return breaker
        case sectionType == analyze.CURLY_SECTION:
            breaker := func(line string)(bool) {
                if line == "}" {
                    return true
                } else {
                    return false
                }
            }
            return breaker
        default:
            breaker := func(line string)(bool) {
                    return false
            }
            return breaker
    }
    //
}


