package cuda

/*
var TARGET_LINE    int = 8000
var TARGET_SECTION int = 8002
var TARGET_FILE    int = 8004
var TARGET_DIR     int = 8008
*/

type Dynima struct {
    // // 
    // //  dynima stores 
    // //
    // each file may got several dynimas binded to itself
    filters   []Filter
    targets   []*Target // ????  seems it is not necessary to store file and directory content inside dynima
    dataSet   []Data    // data will collected while targets processing
}

/*
type Target struct {  // interface {
    // Get       ()(lineAsArray [][]string, err error)
    // GetType ()(typ int)
    // Gather    ()(error)
    // PushPart  ([][]string)(error)
    typ         int
    path        string
    lineAsArray [][]int

}
*/


func(d *Dynima)AppendFilter(f *Filter)(error){
    return nil
}

func(d *Dynima)RunFilters()(r *Result, err error){
    //
    // apply filters targets data
    //
    return r,err
}

func(d *Dynima)SetSource(t *Target)(error){
    // unnecessary 
    return nil
    //
}

func(d *Dynima)getTarget(tgt_id int)(t *Target,err error){
    //
    return
    //
}

func(d *Dynima)getChildTargets(tgt_id int)(tgts *[]Target, err error){
    //
    return
    //
}
/*
func InitiateNewTarget(typ int)(t *Target){
    return t
}*/

