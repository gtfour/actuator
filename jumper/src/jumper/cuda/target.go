package cuda

//import "jumper/common/arrays"


//
var TARGET_UNDEFINED             int    = 7999
var TARGET_UNDEFINED_STR         string = "TARGET_UNDEFINED"

var TARGET_LINE                  int    = 8000
var TARGET_LINE_STR              string = "TARGET_LINE"

var TARGET_SECTION               int    = 8002
var TARGET_SECTION_STR           string = "TARGET_SECTION"

var TARGET_FILE                  int    = 8004
var TARGET_FILE_STR              string = "TARGET_FILE"

var TARGET_DIR                   int    = 8008
var TARGET_DIR_STR               string = "TARGET_DIR"

//
var TARGET_LINE_TYPE_SINGLE      int    = 8101 // whole line placed into first element of lines array
var LINE_SINGLE_STR              string = "SINGLE_LINE"
var TARGET_LINE_TYPE_SPLITTED    int    = 8102 // splitted line placed inside lines array
var LINE_SPLITTED_STR            string = "LINE_SPLITTED"
//


type Target struct {
    //
    //
    //  # Get:       ()(lineAsArray [][]string, err error)
    //  # GetType:   ()(typ int)
    //  # Gather:    ()(error)
    //  # PushPart:  ([][]string)(error)
    //
    //  # Target could  store content of line, file or also just decribe an directory 
    //  # correction: section could not be determined as section on this level of processing
    //
    //
    selfIndex       int         //  // self uniq   number 
    parentIndex     int         //  // uniq parent target number
    typ             int
    path            string
    //lineAsArray   [][]string
    lines           []string
    configured      bool
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

func(t *Target)PushPart(part []string)(err error){
    //  
    //  pushing data to lineAsArray
    //  t.lineAsArray,err = arrays.Extend(t.lineAsArray, part)

    //
    return err
    //
    //
}



/*
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
*/

func InitiateNewTarget(config map[string]string)(t *Target,err error){
    //
    var new_target Target
    //
    target_type, typ_exist  := config["type"]
    target_path, path_exist := config["path"]
    //
    if typ_exist == false { return nil, targetTypeHasNotBeenSpecified }
    if target_type == TARGET_FILE_STR || target_type == TARGET_DIR_STR {
        if path_exist == false { return nil, pathHasNotBeenSpecified } else {
            new_target.path = target_path
            if target_type == TARGET_FILE_STR { new_target.typ = TARGET_FILE } else { new_target.typ = TARGET_DIR }
            new_target.configured = true
            return &new_target, nil
        }
    } else if target_type == TARGET_LINE_STR || target_type == LINE_SINGLE_STR || target_type == LINE_SPLITTED_STR {
        if path_exist == true { return nil,  pathHaveToBeEmpty }
        //line, line_exist := config["line"]
        if target_type == TARGET_LINE_STR || target_type == LINE_SINGLE_STR {
            new_target.typ = TARGET_LINE_TYPE_SINGLE
        } else {
            new_target.typ = TARGET_LINE_TYPE_SPLITTED
        }
        new_target.configured = true
        return &new_target, nil
    }
    //
    // targetTypeHasNotBeenSpecified
    //
    return nil, cantCreateNewTarget
}

func(t *Target)gatherLine(line string)(err error){
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
