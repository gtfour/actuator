package flexi

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
    err       =   notInterfaceSlice
    slice     :=  suspectSlice.([]interface{})
    err       =   notIntSlice
    indexes   :=  suspectListOfIndexes.([]int)
    err       =   nilErr
    //
    //
    for i := 0; i < len(slice); i++ {
        skip := false
        elem := slice[i]
        for z := range indexes {
            indexToRemove := indexes[z]
            if indexToRemove == i {
                // indexToRemove >= 0 && indexToRemove < len(slice)
                // new_slice = append(new_slice[:indexToRemove], new_slice[indexToRemove+1:]...)
                skip = true
                break
            }
        }
        if skip == false {
            new_slice = append(new_slice, elem)
        }
    }
    //
    //
    return
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

//
//
//

func getBySingleIndex( suspectSlice interface{}, suspectIndex interface{} )(new_slice []interface{}, err error) {
    //
    //
    var nilErr error = nil
    defer func() {
        if r := recover(); r != nil {
            // err = notInterfaceSlice
        }
    }()
    err   =  notInterfaceSlice
    slice :=  suspectSlice.([]interface{})
    err   =  notInt
    index := suspectIndex.(int)
    err   =  nilErr
    //
    if index>=0 && index<len(slice) {
        new_slice = append( new_slice, slice[index] )
    }
    //
    return
    //
    //
}




func getByListOfIndexes( suspectSlice interface{}, suspectListOfIndexes interface{} )(new_slice []interface{}, err error) {
    //
    //
    var nilErr error = nil
    defer func() {
        if r := recover(); r != nil {
            // err = notInterfaceSlice
        }
    }()
    err       =   notInterfaceSlice
    slice     :=  suspectSlice.([]interface{})
    err       =   notIntSlice
    indexes   :=  suspectListOfIndexes.([]int)
    err       =   nilErr
    //
    //
    for i := range indexes {
        index := indexes[i]
        if index>=0 && index<len(slice) {
            new_slice = append( new_slice, slice[index] )
        }
    }
    //
    //
    return
}


func GetString(suspectString interface{})(read_string string, err error){
    //
    var nilErr error = nil
    defer func() {
        if r := recover(); r != nil {
        }
    }()
    err         =  notString
    read_string = suspectString.(string)
    err         =  nilErr
    return
    //
}



func Get( slice interface{}, index interface{} )(new_slice []interface{}, err error) {
    //
    new_slice,err = getBySingleIndex(slice, index)
    if err != nil {
        new_slice, err = getByListOfIndexes( slice, index )
        return new_slice, err
    } else {
        return new_slice, err
    }
    //
}



