package cuda

import "jumper/common/arrays"


var TARGET_UNDEFINED int = 7999
var TARGET_LINE      int = 8000
var TARGET_SECTION   int = 8002
var TARGET_FILE      int = 8004
var TARGET_DIR       int = 8008



type Target struct {
    //
    // Get       ()(lineAsArray [][]string, err error)
    // GetType   ()(typ int)
    // Gather    ()(error)
    // PushPart  ([][]string)(error)
    //
    // # Target could store content of line, section, file or also just decribe an directory 
    //
    selfIndex    int // self uniq number 
    parentIndex  int // uniq parent target number
    typ          int
    path         string
    lineAsArray  [][]string
    //
}



func(t *Target)Get()(lineAsArray [][]string,err error) {
    return
}

func(t *Target)GetType()(typ int){
    return
}

func(t *Target)Gather()(err error){

    // var TARGET_LINE    int = 8000
    // var TARGET_SECTION int = 8002
    // var TARGET_FILE    int = 8004
    // var TARGET_DIR     int = 8008

    switch target_type:=t.typ; target_type {
        case TARGET_LINE:
        //case TARGET_SECTION:
        case TARGET_FILE:
        case TARGET_DIR:
    }
    return
}

func(t *Target)PushPart(part [][]string)(err error){
    //  
    // pushing data to lineAsArray
    t.lineAsArray,err = arrays.Extend(t.lineAsArray, part)
    return err
    //
    //
}




func InitiateNewLineTarget(line string)(t *Target){
    //
    return t
    //
}

func InitiateNewFileOrDirectoryTarget(fpath string)(t *Target){
    //
    return t
    //
}

func(t *Target)gatherLine()(err error){
    //
    return err
    //
}

func(t *Target)gatherFile()(err error){
    //
    return err
    //
}

func(t *Target)gatherDir()(err error){
    //
    return err
    //
}
