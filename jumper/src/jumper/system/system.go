package system

import "os"

var TYPE_FILE           int = 3000
var TYPE_DIR            int = 3002
var UNABLE_TO_OPEN_FILE int = 3005
var TYPE_UNDEFINED      int = 3007

func GetFileType(path string)(int){
    f,err := os.Open(path)
    return TYPE_UNDEFINED
}
