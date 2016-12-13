package run

import "os"
import "fmt"
import "strings"


func ParseCmd()(){
    //
    //
    ArgDoubleSplitter(os.Args)
    //
    //
}

func ArgDoubleSplitter(args []string)(arg_pairs [][]string){
    //
    //
    // Split and Glue 
    for i:= range args {
        arg:=args[i]
        fmt.Printf("---\n%s  Is Dup Suffix %v\n---",arg,IsDupSuffix(arg))
    }
    return
    //
    //
}

func IsDupSuffix(arg string)(bool){
    if strings.IndexAny(arg,"--")>=0 {
        return true
    }else {
        return false
    }
}


