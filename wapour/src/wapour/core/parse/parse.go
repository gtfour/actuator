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
        line_slice:=strings.Split(line,"")
        delims,data:=cuda.GetIndexes(line_slice)
        fmt.Printf("-- %v -- %v --\n",delims,data)
    }
    return lines,err
}

func ParseLine(line string)(entry map[string]string) {
    line_slice:=strings.Split(line,"")
    delims,data:=cuda.GetIndexes(line_slice)
    return entry
}

//func GetConfig()()
