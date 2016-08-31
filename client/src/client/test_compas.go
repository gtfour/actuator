package main

import "fmt"
import "client/compas"

func main() {

    var a_list =[][]string {[]string{"a"},[]string{"b"},[]string{"c"}}
    var b_list =[][]string {[]string{"a","yozi"}, []string {"a"} , []string{"b"}, []string{"c"}, []string{"d"}}

    fmt.Printf("\nHello\na_list: %v\nb_list: %v\n",a_list,b_list)
    compas.Trite(a_list, b_list)

    fmt.Printf("\n==Make Mix==\na:\n%v\nb:\n%v",compas.MakeMix(a_list),compas.MakeMix(b_list))

}
