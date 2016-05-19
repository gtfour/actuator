package marconi

type ChangeRequest struct {



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
