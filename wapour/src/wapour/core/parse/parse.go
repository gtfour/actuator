package parse

import "strings"
import "io/ioutil"

func ReadFileLines(filename string) (lines []string,err error){
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return lines, err
    }
    lines = strings.Split(string(content), "\n")
    return lines,err
}
