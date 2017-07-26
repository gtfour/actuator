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

func AppendString(suspectStringSlice interface{}, str string)(stringSlice []interface{},err error){
    //
    //  previous return arguments set: (stringSlice []string,err error)
    //
    fmt.Printf("\n=== >>> Append String : check type %v\n",reflect.TypeOf(suspectStringSlice))
    defer func() {
        if r := recover(); r != nil {
            // err = notStringSlice
            err = notInterfaceSlice
        }
    }()
    // stringSlice = suspectStringSlice.([]string)
    stringSlice = suspectStringSlice.([]interface{})
    stringSlice = append(stringSlice, str)
    return


}
