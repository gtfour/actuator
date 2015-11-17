package main

import "fmt"
import "reflect"
import "errors"
import "client_side/actuator"

var   structFieldCountMissmatched = errors.New("structs have different fields count")

type CompNote struct {

    Attr    string
    Before  string
    After   string


}



func main(){

    test1:=actuator.Prop{DirContentAvailable:true,Inode:67,Perm:"rwx"}
    test2:=actuator.Prop{DirContentAvailable:false,Inode:67,Perm:"r-x"}
    _,_=CompareProp(test1,test2)


    value1:=reflect.ValueOf(test1)
    value2:=reflect.ValueOf(test2)

    fmt.Printf("\nLen of StructOne  %d\n",value1.NumField())
    fmt.Printf("\nLen of StructTwo  %d\n",value2.NumField())


}


func CompareProp(old_prop,new_prop actuator.Prop)(err error,comparison_result []CompNote) {

    valueOld:=reflect.ValueOf(old_prop)
    valueNew:=reflect.ValueOf(new_prop)

    field:=reflect.TypeOf(old_prop)

    old_field_count := valueOld.NumField()
    //new_field_count := valueNew.NumField()

    for i := 0; i <= old_field_count-1; i++  {

        //fmt.Printf("\nField %s is equal before %s -> after: %s\n",field.Field(i).Name,valueOld.Field(i).String(),valueNew.Field(i).String())
        fmt.Printf("\nfield tag: %s\n",field.Field(i).Tag)
        if fmt.Sprint(valueOld.Field(i))!=fmt.Sprint(valueNew.Field(i)) && string(field.Field(i).Tag)!="ignore" {

            fmt.Printf("\nField: %s is different before: %s -> after: %s\n",field.Field(i).Name,fmt.Sprint(valueOld.Field(i)),fmt.Sprint(valueNew.Field(i)))

        }



    }
    return nil,comparison_result

}

