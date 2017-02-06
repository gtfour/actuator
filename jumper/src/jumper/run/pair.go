package run

var PAIR_TYPE_BOOL    int = 3000 // //  example:  --MAKE_NEW_FOLDER
var PAIR_TYPE_STRING  int = 3002 // //  example:  --dir="/exports"

var YES bool = true
var NO  bool = false

type Pair struct {
    typ                int
    key,value          string
    value_initialized  bool
}

func (p *Pair)IsIncomplete()(bool){
    typ               := p.typ
    key               := p.key
    value             := p.value
    value_initialized := p.value_initialized
    if typ == PAIR_TYPE_BOOL && key != "" {
        return false
    } else if ((typ == PAIR_TYPE_STRING && key != "" && value!="")||(typ == PAIR_TYPE_STRING && key != "" && value_initialized == false)){
        return false
    } else {
        return true
    }
}
