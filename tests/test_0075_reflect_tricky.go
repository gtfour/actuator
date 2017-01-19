package main

import "fmt"
import "reflect"

func main(){

    my_int_var      := 6
    my_bool_var     := true
    my_string_var   := "a"


    fmt.Printf("\nstage:\treflect:\t:\n%d %v %s\n",my_int_var,my_bool_var,my_string_var)
    fmt.Printf("\nvar types: %v  %v  %v\n",reflect.TypeOf(my_int_var), reflect.TypeOf(my_bool_var), reflect.TypeOf(my_string_var))



    myint           := reflect.ValueOf(my_int_var)
    myinttype       := reflect.TypeOf(my_int_var)
    // //   myIntArrayType  := reflect.ArrayOf(1,reflect.TypeOf(my_int_var))
    intSlice        := reflect.MakeSlice(reflect.SliceOf(myinttype),0,1)
    // //   intSlicePtr     := reflect.New(reflect.TypeOf(intSlice))
    // //   intSlicePtr.Set(intSlice)
    // //   intArPtr.Index()
    // //   reflect.Append(intSlicePtr, myint)
    intSlice=reflect.Append(intSlice, myint)
    // //   intSlice[0]     =  my_int_var
    intAr           := intSlice.Interface()
    // //   intAr         =  append( intAr, 2 )

    fmt.Printf("\n|My int interface: %v |My slice interface: %v|Type of interface: %v| Type of slice: %v|\n ", intAr,intSlice, reflect.TypeOf(intAr),reflect.TypeOf(intSlice))

}
