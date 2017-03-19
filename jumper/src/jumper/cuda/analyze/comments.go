package analyze

import "strings"

var comments                  =  []string {`//` , `#`}

func IsComment(entry string) (comment bool) {

    for i:= range comments {
        if strings.HasPrefix(entry, comments[i]) == true {
            return true
        }
    }
    return false
}
