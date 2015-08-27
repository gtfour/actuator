package main

import "fmt"
import "strings"



func main() {

  var test1 []string

  var test2 []string

  test1=[]string{"a","b","c"}
  test2=[]string{"1","2","3","4","5"}
  new_string:=strings.Join(test2[1:]," ")

  fmt.Printf("%s",test1)
  fmt.Printf("%s",test2)
  fmt.Printf("new string %s",new_string)

}
