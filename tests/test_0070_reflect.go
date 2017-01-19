package main

import "fmt"
import "reflect"

// testing MakeSlice function from reflect package
// TypeOf(i interface{}) Type

func main(){
    //fmt.Printf("\nRun the fucking function:\n%v\n",get_the_fucking_type(3))
    three                  := 3
    suspect_int_type       := reflect.TypeOf(three)
    fmt.Printf(">>Type of three:%v\n",suspect_int_type)
    my_new_slice           := reflect.MakeSlice(reflect.SliceOf(suspect_int_type),0,0)
    fmt.Printf(">>Type of new slice:%v\n",my_new_slice.Type())
    my_new_array_value     := reflect.New(my_new_slice.Type())

    // //my_new_int_array       := new(my_new_slice.Type())
    my_new_array_value.Elem().Set(my_new_slice)
    //my_new_array           := my_new_array_value.Elem()
    my_new_array           =  append(my_new_array_value.Elem(), three)
    fmt.Printf("\nMy int array :%v\n",my_new_array)
}

func get_the_fucking_type(i interface{})(reflect.Type){
    return reflect.TypeOf(i)
}
