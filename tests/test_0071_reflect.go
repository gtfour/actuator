package main

import (
    "fmt"
    "reflect"
)

func main() {
    // one way is to have a value of the type you want already
    // var my_int_array []int
    a := 1
    // reflect.New works kind of like the built-in function new
    // We'll get a reflected pointer to a new int value
    intPtr      := reflect.New(reflect.TypeOf(a))
    intArrayPtr := reflect.New(reflect.ArrayOf(0, reflect.TypeOf(a)))
    // Just to prove it
    //  b := intPtr.Elem().Interface().(reflect.TypeOf(a))

    //  b    := intPtr.Elem().Interface()
    //  b_ar := intArrayPtr.Elem().Interface().(intArrayPtr)

    (*intArrayPtr) = append((*intArrayPtr), intPtr)
    //b_array:=reflect.ArrayOf(0, intPtr)
    // // b := intPtr.Elem().Interface().(reflect.PtrTo(reflect.TypeOf(a)))
    // Prints 0
    fmt.Println(b)

    // We can also use reflect.New without having a value of the type
    var nilInt *int
    intType := reflect.TypeOf(nilInt).Elem()
    intPtr2 := reflect.New(intType)
    // Same as above
    c := intPtr2.Elem().Interface().(int)
    // Prints 0 again
    fmt.Println(c)
}
