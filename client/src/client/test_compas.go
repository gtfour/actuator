package main

import "fmt"
import "client/compas"

func main() {

    var a_list =[][]string {[]string{"a"},[]string{"b"},[]string{"c"}}
    var b_list =[][]string {[]string{"a"}, []string{"b"}, []string{"c"}, []string{"d"}}

    fmt.Printf("\nHello\na_list: %v\nb_list: %v\n",a_list,b_list)
    compas.Trite(a_list, b_list)

}
