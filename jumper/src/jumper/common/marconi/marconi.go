package marconi

import "encoding/json"

var STATUS_OK   = 8000
var STATUS_FAIL = 8001


type Request struct {
    ChangeType   int
    ObjName      string
    ObjType      string
    ObjPath      string
}

type DataUpdate struct {

    SourceType  string // file or command ( actuator or balckout  )
    SourceName  string
    SourcePath  string // /filename or /command_name
    UpdateType  string // Update,Append,Remove,RemoveFile
    UpdateData  string // 
    DataHash    string
    ServerTime  string
    ServerId    string

}

type Response struct {
    Status int
}


func (r *Response) GetRaw ()([]byte, error) {
    raw,err:=json.Marshal(r)
    return raw,err
}

func (r *Request) GetRaw ()([]byte, error) {
    raw,err:=json.Marshal(r)
    return raw,err
}

