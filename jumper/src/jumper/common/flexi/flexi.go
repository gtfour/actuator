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




