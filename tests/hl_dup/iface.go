package main

import "fmt"
import "strings"
import "io/ioutil"

type Log struct {

    offset_steps [][2]int // step size is  100

}

type Entry struct {

}

func (log *Log)IncreaseOffsetsByOne(since int)(){



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

func Parse(path string)(error){
    return nil
}

func (file *Log)Insert(entry string)(error){
    return nil
}

func SplitIp(ip_addr string)(ip_slice []int,err error){
    ip:=strings.Split(ip_addr,".")
    fmt.Printf("\nip address: %v\n",ip)
    return ip_slice,err
}
