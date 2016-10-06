package parse

import "fmt"
import "strings"
import "io/ioutil"

func ReadFileLines(filename string) (lines []string,err error){
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return lines, err
    }
    lines = strings.Split(string(content), "\n")
    for i := range lines {
        line:=lines[i]
        line_slice:=strings.Split(line,"")
        delims,data:=GetIndexes(line_slice)
        fmt.Printf("-- %v -- %v --\n",delims,data)
    }
    return lines,err
}

func ParseLine(line string)(entry map[string]string) {
    return entry
}

//func GetConfig()()
