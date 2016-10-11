package cuda

import "client/cross"

type Clutch struct {
    // middleware between cross.Dynima and parsers
    filter_sequence []string
}


func CreateClutch(d *cross.Dynima)(c Clutch){
    return c
}
