package cuda


var TARGET_LINE    int = 8000
var TARGET_SECTION int = 8002
var TARGET_FILE    int = 8004
var TARGET_DIR     int = 8008



type Target struct {
    //
    //  Get       ()(lineAsArray [][]string, err error)
    //  GetType   ()(typ int)
    //  Gather    ()(error)
    //  PushPart  ([][]string)(error)
    //
    typ         int
    path        string
    lineAsArray [][]int
}



func(t *Target)Get()(lineAsArray [][]string, err error) {
    return
}

func(t *Target)GetType()(typ int) {
    return
}

func(t *Target)Gather()(err error) {
    return
}

func(t *Target)PushPart([][]string)(err error){
    return
}




func InitiateNewTarget(typ int)(t *Target){
    return t
}
