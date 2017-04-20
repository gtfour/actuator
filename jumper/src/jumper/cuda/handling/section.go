package handling

import "jumper/cuda/analyze"

func GetSectionBreaker( line string, name [2]int, tag [2]int, sectionType int )(func(string)(bool)){
    //
    // breaker will stop filling Lines to the current section
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
                match := "</" + line[name[0]:name[1]+1]+">"
                if phrase == match { return true } else { return false }
            }
            return breaker
        case sectionType == analyze.CURLY_SECTION:
            breaker := func(line string)(bool){
                if line == "}" {
                    return true
                } else {
                    return false
                }
            }
            return breaker
        default:
            breaker := func(line string)(bool){
                    return false
            }
            return breaker
    }
    //
}


