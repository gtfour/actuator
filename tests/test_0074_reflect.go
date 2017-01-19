package main

import "fmt"
import "reflect"

func main(){

    my_int_var      := 0
    my_bool_var     := true
    my_string_var   := "a"


    fmt.Printf("\nstage:\treflect:\t:\n%d %v %s\n",my_int_var,my_bool_var,my_string_var)
    fmt.Printf("\nvar types: %v  %v  %v\n",reflect.TypeOf(my_int_var), reflect.TypeOf(my_bool_var), reflect.TypeOf(my_string_var))


    myIntArrayType  := reflect.ArrayOf(0,reflect.TypeOf(my_int_var))
    intArPtr        := reflect.New(myIntArrayType)
    intAr           := intArPtr.Interface()
    
    //intAr         =  append( intAr, 2 )

    fmt.Printf("\nMy int array as interface : %v     Type : %v  \n ", intAr, reflect.TypeOf(int))

}
