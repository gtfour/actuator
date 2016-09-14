package main

import "strings"
import "io/ioutil"

type Log struct {

}

type Entry struct {

}




func ReadFileLines(filename string) (lines []string,err error){
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        return lines, err
    }
    lines = strings.Split(string(content), "\n")
    return lines,err
}



func OpenFile(path string)(error){
    return nil
}

func Parse(path string)(){
    return nil
}

func (file *Log)Insert(entry string)(error){
    return nil
}
