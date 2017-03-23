package cuda

import "sync"
import "jumper/cuda/targets"
import "jumper/cuda/filtering"

/*

//  var  TARGET_LINE    int = 8000
//  var  TARGET_SECTION int = 8002
//  var  TARGET_FILE    int = 8004
//  var  TARGET_DIR     int = 8008

*/

type Mark struct {
    // bounding on change of an  directory/file or change of an command output

}

type Dynima struct {
    //
    //  ::
    //  ::
    //  ::  dynima stores  
    //  ::  each file may got several dynimas binded to itself
    //  ::
    //
    sync.RWMutex                             //  mutex will be used to freze operations over dynima while changing filters or modifying targets
    filters          filtering.FilterList    // 
    targets          targets.TargetList      //  ????  seems it is not necessary to store file and directory content inside dynima
    configured       bool                    //
    offset           int64                   //  for log files
    mark             Mark
    //
    //  dataSet  []Data                      // data will collected while targets processing
    //  ::
    //  ::
    //  ::
    //  ::
    //  ::
    //
}

/*
type Target struct {  // interface {
    //
    //  Get       ()(lineAsArray [][]string, err error)
    //  GetType ()(typ int)
    //  Gather    ()(error)
    //  PushPart  ([][]string)(error)
    //
    typ         int
    path        string
    lineAsArray [][]int

}
*/


func(d *Dynima)RunFilters()(r *Result, err error){
    //
    // apply filters targets data
    //
    d.Lock()
    defer d.Unlock()
    //
    return r,err
    //
    //
    //
}

func(d *Dynima)AppendFilter(f *filtering.Filter)(error){
    //
    //
    return nil
    //
    //
}

func(d *Dynima)SetSource(t *targets.Target)(error){
    // unnecessary 
    //
    return nil
    //
    //
}

func(d *Dynima)AppendTarget(t *targets.Target)(error){
    // 
    //
    return nil
    //
    //
}
//
//
func(d *Dynima)RemoveTarget(t *targets.Target)( error ){
    // 
    //
    return nil
    //
    //
}
//
//
func(d *Dynima)getTarget(tgt_id int)( t *targets.Target, err error ){
    //
    //
    return
    //
    //
}
//
//
func(d *Dynima)getChildTargets(parent_target_id int)(child_targets *[]targets.Target, err error){
    //
    //
    return
    //
    //
}
//
//
func NewDynima()( *Dynima ){
    //
    //
    var d Dynima
    d.filters     = make( filtering.FilterList, 0 )
    d.targets     = make( targets.TargetList,   0 )
    return &d
    //
    //
}
//
//
