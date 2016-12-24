package run

import "os"
import "fmt"
import "strings"
import "jumper/cuda"

func GetFilterSet()(d *cuda.Dynima){


}


func ParseCmd()(){
    //
    //
    InputArgSplitter(os.Args)
    //
    //
}

func InputArgSplitter(args []string)(arg_pairs [][]string){
    //
    //
    // Split and Glue 
    //previous_arg_is_bool := false // Is true if currect arg is flag and previous arg is also flag
    for i:= range args {
        arg                        := args[i]
        isDupSuffix,isSingleSuffix := CheckSuffix(arg)
        fmt.Printf("---\n%s|Is Dup Suffix %v|Is Single Suffix %v|\n---",arg,isDupSuffix,isSingleSuffix)
    }
    return
    //
    //
}

func IsDupSuffix(arg string)(bool){
    if strings.HasPrefix(arg,"--"){
        return true
    }else {
        return false
    }
}

func IsSingleSuffix(arg string)(bool){
    if strings.HasPrefix(arg,"-"){
        return true
    }else {
        return false
    }
}

func CheckSuffix(arg string)(bool,bool){
    dup    := IsDupSuffix(arg)
    single := IsSingleSuffix(arg)
    if dup { single = false }
    return dup,single
}

func KeySplitter(arg string)(values [2]string){
    return values
}

