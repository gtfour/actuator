package main

import (
    "fmt"
    "reflect"
)

func main() {
    // one way is to have a value of the type you want already
    a := 1
    var int_array []int
    // reflect.New works kind of like the built-in function new
    // We'll get a reflected pointer to a new int value
    intPtr := reflect.New(reflect.TypeOf(a))
    // Just to prove it
    b := intPtr.Elem().Interface().(reflect.TypeOf(a).Elem())
    // Prints 0
    int_array = append(int_array, b)
    fmt.Printf("b:%d\nint_array:%v\n",b,int_array)

    // We can also use reflect.New without having a value of the type
    var nilInt *int
    intType := reflect.TypeOf(nilInt).Elem()
    intPtr2 := reflect.New(intType)
    // Same as above
    c := intPtr2.Elem().Interface().(int)
    // Prints 0 again
    fmt.Println(c)
}
