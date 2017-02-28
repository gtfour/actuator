package main


import "fmt"
import "jumper/common/arrays"

func main(){

    my_piece          := []string{ "x", "y" , "z" }
    my_part           := [][]string{my_piece}

    text              := []string{ "a","b","c" }
    my_array          := [][]string{ text }

    my_new_array,_  := arrays.Extend(my_array, my_part)
    my_new_array,_  =  arrays.Extend(my_new_array, my_part)
    my_new_array,_  =  arrays.Extend(my_new_array, my_part)

    fmt.Printf("\n<< Extended array: %v\n",my_new_array)


}
