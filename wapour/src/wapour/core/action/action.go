package action

var ATYPE_MODAL int8 =  0
var ATYPE_LINK  int8 =  1

type Action struct {

    Name        string
    Type        int8
    SuccessUrl  string
    PostUrl     string

}
