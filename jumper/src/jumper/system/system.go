package system

var TYPE_FILE      int = 3000
var TYPE_DIR       int = 3002
var TYPE_UNDEFINED int = 3005

func GetFileType(path string)(int){
    return TYPE_UNDEFINED
}
