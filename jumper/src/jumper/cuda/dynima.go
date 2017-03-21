package cuda

import "sync"
import "jumper/cuda/filtering"

/*
var TARGET_LINE    int = 8000
var TARGET_SECTION int = 8002
var TARGET_FILE    int = 8004
var TARGET_DIR     int = 8008
*/

type Dynima struct {
    //
    // :
    // :  dynima stores  
    // :  each file may got several dynimas binded to itself
    // :
    //
    sync.RWMutex             // mutex will be used to freze operations over dynima while changing filters or modifying targets
    filters        filtering.FilterList // 
    targets        []*Target  // ????  seems it is not necessary to store file and directory content inside dynima
    //
    // dataSet  []Data       // data will collected while targets processing
    //
    // :
    // :
    // :
    // :
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


func(d *Dynima)AppendFilter(f *filtering.Filter)(error){
    //
    return nil
    //
}
//
func(d *Dynima)RunFilters()( r *Result, err error ){
    //
    // apply filters targets data
    //
    d.Lock()
    defer d.Unlock()
    //
    return r,err
    //
    //
}
//
func(d *Dynima)SetSource(t *Target)(error){
    // unnecessary 
    //
    return nil
    //
    //
}
//
func(d *Dynima)AddTarget(t *Target)(error){
    // 
    //
    return nil
    //
    //
}
//
func(d *Dynima)RemoveTarget(t *Target)( error ){
    // 
    //
    return nil
    //
    //
}
//
func(d *Dynima)getTarget(tgt_id int)( t *Target, err error ){
    //
    //
    return
    //
    //
}
//
func(d *Dynima)getChildTargets(parent_target_id int)(child_targets *[]Target, err error){
    //
    //
    return
    //
    //
}
//
