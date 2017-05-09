package run

import "os"
import "fmt"
import "strings"
import "jumper/cuda"

func GetFilterSet()(d *cuda.Dynima){
    return
}

func ParseCmd()(){
    //
    InputArgSplitter(os.Args)
    //
}

func InputArgSplitter(args []string)(arg_pairs []Pair){
    //
    // Split and Glue 
    // previous_arg_is_bool := false // Is true if currect arg is flag and previous arg is also flag
    //
    prev_is_hyphenized := false
    this_is_hyphenized := false
    this_is_value      := false
    beginning          := true
    // pair_created    := false
    // apair           := Pair{}
    for i := range args {
        if i != 0 { beginning = false }
        arg                                :=  args[i]
        is_dup_hyphen, is_single_hyphen    :=  CheckHyphen(arg)
        this_is_hyphenized                 =   is_dup_hyphen || is_single_hyphen
        this_is_value                      =   !this_is_hyphenized
        //
        // finish check
        //
        prev_is_hyphenized                  =  this_is_hyphenized
        fmt.Printf("---\n%s---|Is Dup Suffix %v|Is Single Suffix %v|Is Value %v|\n---", arg, is_dup_hyphen, is_single_hyphen,this_is_value )
    }
    _ = prev_is_hyphenized
    _ = beginning
    return
    //
}

func IsDupHyphen(arg string)(bool){
    if strings.HasPrefix(arg, "--"){
        return true
    } else {
        return false
    }
}

func IsSingleHyphen(arg string)(bool){
    if strings.HasPrefix(arg, "-"){
        return true
    } else {
        return false
    }
}

func CheckHyphen(arg string)(bool,bool){
    dup    := IsDupHyphen(arg)
    single := IsSingleHyphen(arg)
    if dup { single = false }
    return dup,single
}

func KeySplitter(arg string)(values [2]string){
    return values
}

func NextElemExist(ci int,array_len int)( ex bool){
    if ci < array_len { ex = true }
    return
}
