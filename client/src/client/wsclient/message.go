package wsclient

import "encoding/json"

// marconi message

type Message struct {
    DataType   string `json:"datatype"`
    Data       json.RawMessage `json:"data"`
}

type DataUpdate struct {

    SourceType string // file or command ( actuator or blackout  )
    SourceName string
    SourcePath string // /filename or /command_name
    UpdateType string // Update,Append,Remove,RemoveFile
    UpdateData string //
    DataHash   string
    ServerTime string
    ServerId   string

}

func ( message *Message) GetRaw ()([]byte, error) {
    raw,err:=json.Marshal(message)
    return raw,err
}

func ( data *DataUpdate) GetRaw ()([]byte, error) {
    raw,err:=json.Marshal(data)
    return raw,err
}

type DataModify struct {

    DynimaId   string
    SourceType string // file or command ( actuator or blackout  )
    SourceName string
    SourcePath string // /filename or /command_name
    UpdateType string // Update,Append,Remove,RemoveFile
    UpdateData string //
    DataHash   string
    ServerTime string
    ServerId   string

}
