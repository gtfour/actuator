package evebridge

import "time"
import "fmt"
import "client/wsclient"

var LOG_CHANNEL_TIMEOUT_MS  time.Duration  = 1000

const (
      INITIALIZED  =  0 // initialized
      CREATED      =  1
      MODIFIED     =  2
      REMOVED      =  3)

type Event struct {

    Date string
    Path string
    Type string

}

type CompNotes struct {
    Path  string
    State int8
    List  []CompNote
}

type CompNote struct {
    Field    string
    Before   string
    After    string
}
type MngMessage struct {
    action string
    path   string
}

/*
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
*/

func Handle(messages chan CompNotes )() {
        var websocket_connection = wsclient.WsConn
        for {
            select{
                case message:=<-messages:
                    var ws_message_data = wsclient.DataUpdate{SourcePath:message.Path}
                    var ws_message = &wsclient.Message{DataType:"data_update",Data:ws_message_data}
                    fmt.Printf("\nStart writing\n")
                    websocket_connection.Write(ws_message)
                    fmt.Printf("\nFinish writing\n")
                    fmt.Printf("Message: %v HaveToParse: %v\n",message,message.FieldExists("HashSum"))
                default:
                    time.Sleep( LOG_CHANNEL_TIMEOUT_MS  * time.Millisecond )
                    //fmt.Println("No messages")
            }
        }
}




func (cn *CompNotes) FieldExists ( field string )(exists bool) {

    for cnote_id := range cn.List {
        cnote:=cn.List[cnote_id]
        if cnote.Field == field { exists = true ; break  }
    }
    return exists
}
