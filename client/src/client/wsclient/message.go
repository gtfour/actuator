package wsclient

// marconi message

type MessageUpdate struct {
    DataType   string          `json:"datatype"`
    Data       DataUpdate      `json:"data"`
}

type DataUpdate struct {

    SourceType string // file or command ( actuator or balckout  )
    SourceName string
    SourcePath string // /filename or /command_name
    UpdateType string // Update,Append,Remove,RemoveFile
    UpdateData string //
    DataHash   string
    ServerTime string
    ServerId   string

}




