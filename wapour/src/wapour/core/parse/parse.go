package parse

import "fmt"
import "strings"
import "io/ioutil"
import "jumper/cuda"

func ReadFileLines(filename string) (lines []string,err error){
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return lines, err
    }
    lines = strings.Split(string(content), "\n")
    for i := range lines {
        line:=lines[i]
        _=ParseLine(line)
    }
    return lines,err
}

func ParseLine(line string)(entry map[string]string) {
    line_slice:=strings.Split(line,"")
    delims,data:=cuda.GetIndexes(line_slice)
    delims,data=cuda.QuotesFilter(line_slice,delims,data)
    cuda.DebugPrintCharCounter(line)
    fmt.Printf("\ndelims: %v\n data: %v\n" , delims , data)
    return entry
}

//func GetConfig()()
