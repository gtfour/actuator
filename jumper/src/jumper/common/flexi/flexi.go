package flexi

import "fmt"
import "reflect"

func GetTheFuckingArray(i interface{})(interface{}){

    // return interface containing array  with same type like i's
    // WARNING: was not tested

    inVar      := reflect.ValueOf(i)
    inVarType  := reflect.TypeOf(inVar)
    inSlice    := reflect.MakeSlice(inVarType,0,1)

    inSlice    =  reflect.Append(inSlice, inVar)
    return inSlice.Interface()
}

func appendInterface(suspectInterfaceSlice interface{}, suspectInterfaceToAppend interface{})(interfaceSlice []interface{},err error){
    defer func() {
        if r := recover(); r != nil {
            err = notInterfaceSlice
        }
    }()
    interfaceSlice    =  suspectInterfaceSlice.([]interface{})
    interfaceToAppend := suspectInterfaceToAppend.(interface{})
    interfaceSlice = append(interfaceSlice, interfaceToAppend)
    return
}


func appendInterfaceFrom(suspectInterfaceSlice interface{}, suspectSourceInterfaceSlice interface{})(extendedInterfaceSlice []interface{},err error) {

    defer func() {
        if r := recover(); r != nil {
            err = notInterfaceSlice
        }
    }()
    interfaceSlice       := suspectInterfaceSlice.([]interface{})
    sourceInterfaceSlice := suspectSourceInterfaceSlice.([]interface{})
    for i := range sourceInterfaceSlice {
        interfaceToAppend := sourceInterfaceSlice[i]
        interfaceSlice    = append(interfaceSlice, interfaceToAppend)
    }
    return interfaceSlice, err
}

func Append( slice interface{}, new_value interface{} )(new_slice []interface{}, err error) {
    //
    new_slice, err = appendInterfaceFrom( slice, new_value )
    if err != nil {
        new_slice, err = appendInterface( slice, new_value )
        return new_slice, err
    } else {
        return new_slice, err
    }
    //
}

func removeBySingleIndex( suspectSlice interface{}, suspectIndex interface{} )(new_slice []interface{}, err error) {
    //
    //
    var nilErr error = nil
    defer func() {
        if r := recover(); r != nil {
            // err = notInterfaceSlice
        }
    }()
    err       =  notInterfaceSlice
    new_slice =  suspectSlice.([]interface{})
    err       =  notInt
    index     := suspectIndex.(int)
    err       =  nilErr
    //
    new_slice = append(new_slice[:index], new_slice[index+1:]...)
    //
    return
    //
    //
}


func removeByListOfIndexes( suspectSlice interface{}, suspectListOfIndexes interface{} )(new_slice []interface{}, err error) {
    //
    //
    var nilErr error = nil
    defer func() {
        if r := recover(); r != nil {
            // err = notInterfaceSlice
        }
    }()
    err       =  notInterfaceSlice
    new_slice =  suspectSlice.([]interface{})
    err       =  notIntSlice
    indexes   := suspectListOfIndexes.([]int)
    err       =  nilErr
    fmt.Printf("flexi:checking len of input slice : %d",len(new_slice))
    for i := range indexes {
        indexToRemove := indexes[i]
        if indexToRemove >= 0 && indexToRemove < len(new_slice) {
            new_slice = append(new_slice[:indexToRemove], new_slice[indexToRemove+1:]...)
        }
    }
    //
    return
    //
    //
}


func Remove( slice interface{}, index interface{} )(new_slice []interface{}, err error) {
    //
    new_slice,err = removeBySingleIndex(slice, index)
    if err != nil {
        new_slice, err = removeByListOfIndexes( slice, index )
        return new_slice, err
    } else {
        return new_slice, err
    }
    //
}
