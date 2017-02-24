package cuda

/*
var TARGET_LINE    int = 8000
var TARGET_SECTION int = 8002
var TARGET_FILE    int = 8004
var TARGET_DIR     int = 8008
*/

type Dynima struct {
    filters  []Filter
    target   *Target
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
    return r,err
}

func(d *Dynima)SetSource(t *Target)(error){
    return nil
}

/*
func InitiateNewTarget(typ int)(t *Target){
    return t
}*/

