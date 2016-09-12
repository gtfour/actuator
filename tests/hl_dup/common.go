package main

import "strings"
import "io/ioutil"

type MyFile struct {



}

func ReadFileLines (filename string) (lines []string,err error){


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

func (file *MyFile)AddEntry()(){

}
