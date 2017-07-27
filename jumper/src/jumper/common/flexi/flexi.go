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

func AppendInterface(suspectInterfaceSlice interface{}, suspectInterfaceToAppend interface{})(interfaceSlice []interface{},err error){
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


func AppendInterfaceFrom(suspectInterfaceSlice interface{}, suspectSourceInterfaceSlice interface{})(extendedInterfaceSlice []interface{},err error) {

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


