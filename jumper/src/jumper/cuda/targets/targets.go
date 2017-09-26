package targets

//import "jumper/common/arrays"
import "path"
import "strconv"
import "jumper/common/file"
import "jumper/common/flexi"

type TargetListPtrs  []*Target
type TargetList      []Target

type Target struct {
    //
    //
    //  #  Get:       ()(lineAsArray [][]string, err error)
    //  #  GetType:   ()(typ int)
    //  #  Gather:    ()(error)
    //  #  PushPart:  ([][]string)(error)
    //
    //  #  Target could  store content of line, file or also just decribe an directory 
    //  #  correction: section could not be determined as section on this level of processing
    //
    //
    selfIndex                 int          //  // self uniq   number 
    parentIndex               int          //  // uniq parent target number
    typ                       int
    path                      string
    pathShort                 string
    // lineAsArray            [][]string
    lines                     []string
    configured                bool
    gatherFailed              bool
    //
    diving                    bool         //  // gathering nested directories. seems that i can't implement this feauture yet here
    nestedTargets             TargetList   //  // instead of TargetListPtrs
    //
    //
    isLogFile                 bool
    isDirectoryWithLogFiles   bool
    child                     bool
    offset                    int64        // // for log files 
    //
    // # have to add fileds to check command output
    //
}

func(tl *TargetListPtrs)Append(t *Target)(err error){
    if t.configured {
        (*tl) = append((*tl), t)
        return nil
    } else {
        return targetWasNotConfigured
    }
}

func(tl *TargetListPtrs)GetCopy()(targets TargetList){
    for i:= range (*tl){
        t := (*tl)[i]
        var target Target
        target = *t
        targets = append(targets, target)
    }
    return
}


func(tl *TargetList)Append(t *Target)(err error){
    //
    //
    if t.configured {
        //
        var target Target
        target = *t
        (*tl) = append((*tl), target)
        return nil
        //
    } else {
        return targetWasNotConfigured
    }
}

func(tl *TargetListPtrs)IsEmpty()(bool){
    if len(*tl) <= 0 { return true } else { return false }
}

func(tl *TargetList)IsEmpty()(bool){
    if len(*tl) <= 0 { return true } else { return false }
}




func(t *Target)Get()(lineAsArray [][]string,err error) {
    return
}

func(t *Target)GetType()(int)         { return t.typ        }
func(t *Target)GetPath()(string)      { return t.path       }
func(t *Target)GetPathShort()(string) { return t.pathShort  }
func(t *Target)GetLines()([]string)   { return t.lines      }

func(t *Target)PushPart(part []string)(err error){
    //  
    //  pushing data to lineAsArray
    //  t.lineAsArray,err = arrays.Extend(t.lineAsArray, part)

    //
    return err
    //
    //
}


func NewTarget(config map[string]string)(t *Target,err error){
    //
    var new_target Target
    //
    target_type, typ_exist  := config["type"]
    //
    target_path, path_exist := config["path"]
    //
    if typ_exist == false { return nil, targetTypeHasNotBeenSpecified }

    index,  index_exist := config["index"]

    if index_exist {
        index_int,err:=strconv.Atoi(index)
        if err == nil { new_target.selfIndex = index_int }
    }

    parent_index, parent_index_exist := config["parent_index"]
    if parent_index_exist {
        parent_index_int,err:=strconv.Atoi(parent_index)
        if err == nil { new_target.parentIndex = parent_index_int }
    }

    if target_type == TARGET_FILE_STR || target_type == TARGET_DIR_STR || target_type == TARGET_DIRECTORY_STR {
        if path_exist == false { return nil, pathHasNotBeenSpecified } else {
            new_target.path = target_path
            if target_type == TARGET_FILE_STR {
                new_target.typ = TARGET_FILE
            } else {
                new_target.typ       =  TARGET_DIR
                diving, diving_exist := config["diving"]
                if diving_exist {
                    if diving == TRUE  {  new_target.diving = true  }
                    if diving == FALSE {  new_target.diving = false }
                }
            }
            new_target.configured = true
            return &new_target, nil
        }
    } else if target_type == TARGET_LINE_STR || target_type == LINE_SINGLE_STR || target_type == LINE_SPLITTED_STR {
        if path_exist == true { return nil,  pathHaveToBeEmpty }
        //line, line_exist := config["line"]
        if target_type == TARGET_LINE_STR || target_type == LINE_SINGLE_STR {
            new_target.typ   =  TARGET_LINE_TYPE_SINGLE
            line, line_exist := config["line"]
            if line_exist {
                new_line,err     := flexi.GetString(line)
                if err == nil {
                    new_target.lines =  append(new_target.lines, new_line)
                } else {
                    return nil, err
                }
            }
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



func(t *Target)Gather()(err error){

    // var TARGET_LINE    int = 8000
    // var TARGET_SECTION int = 8002
    // var TARGET_FILE    int = 8004
    // var TARGET_DIR     int = 8008
    if !t.configured { return targetWasNotConfigured }
    switch target_type:=t.typ; target_type {
        case TARGET_LINE:
            //
            err = nil
            //
        case TARGET_FILE:
            //
            err = t.gatherFile()
            //
        case TARGET_DIR:
            //
            err = t.gatherDir()
            //
    }
    if err != nil {
        t.gatherFailed = true
    }
    return
}

func (t *Target)SetLines(lines []string)(err error){
    if !t.configured { return targetWasNotConfigured }
    if t.typ == TARGET_LINE_TYPE_SINGLE || t.typ == TARGET_LINE_TYPE_SPLITTED { } else { return cantAddLineForThisTypeOfTarget }
    if lines != nil {
         t.lines = lines
         return nil
    } else {
        return lineIsNil
    }
}

func (t *Target)SetLine(line string)(err error){
    if !t.configured { return targetWasNotConfigured }
    if t.typ != TARGET_LINE_TYPE_SINGLE { return cantAddLineForThisTypeOfTarget }
    t.lines = make([]string,0)
    t.lines = append(t.lines, line)
    return nil
}



func(t *Target)gatherFile()(err error){
    //
    lines,err := file.ReadFile(t.path)
    //
    //
    //  target_config          := make(map[string]string,0)
    //  target_config["type"]  =  "SINGLE_LINE"
    //  tgt,err                := cuda.NewTarget( target_config )
    //
    //
    if err == nil {
        t.lines       = lines
        t.pathShort   = path.Base(t.path)
    }
    return err
    //
}

func(t *Target)gatherDir()(err error) {
    //
    dir_files,err := file.ReadDirFiles(t.path)
    if err !=nil { return }
    //
    for i := range dir_files {
        //
        //
        dir_file                   :=  dir_files[i]
        targetFileConfig           :=  make(map[string]string,0)
        targetFileConfig["type"]   =   "TARGET_FILE"
        targetFileConfig["path"]   =   dir_file
        tgtFilePtr,err             :=  NewTarget(targetFileConfig)
        if err != nil || tgtFilePtr.configured == false { continue }
        err = tgtFilePtr.Gather()
        //
        //
        if err == nil {
            tgtFilePtr.parentIndex    =  t.selfIndex
            tgtFilePtr.child          =  true
            // var  tgtFile  Target
            // tgtFile                   =  *tgtFilePtr
            // t.nestedTargets           =  append( t.nestedTargets, tgtFile )
            t.nestedTargets.Append( tgtFilePtr )
        }
        //
        //
    }
    //
    t.pathShort = path.Base( t.path )
    //
    return nil
    //
    //
    //
}

func(t *Target)GetNestedTargets()([]Target) {
    /*for i:= range t.nestedTargets {
        nestedTargetAddr:=t.nestedTargets[i]
        var target Target
    }*/
    return t.nestedTargets
}



func(t *Target)CleanLines()() {
    /*for i:= range t.nestedTargets {
        nestedTargetAddr:=t.nestedTargets[i]
        var target Target
    }*/
    t.lines = []string{}
}

func(t *Target)GatherIsFailed()(bool) {   return t.gatherFailed  }
func(t *Target)IsConfigured()(bool)   {   return t.configured    }
