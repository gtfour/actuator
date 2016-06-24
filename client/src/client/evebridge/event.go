package evebridge

import "fmt"
import "time"
import "encoding/json"
import "client/wsclient"
import "client/activa"

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
    //
    SourceType string // file or directory or command ( source: actuator or blackout  )
    SourceName string
    SourcePath string // /filename or /command_name
    //UpdateType string // Update,Append,Remove,RemoveFile
    //UpdateData string //
    DataHash   string
    //ServerTime string
    //ServerId   string
    //
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
        motions := make(chan *activa.Motion, 100)
        var websocket_connection = wsclient.WsConn
        for {
            select{
                case message:=<-messages:
                    fmt.Printf("\n<<Evebridge: message has been recieved>>\n")
                    var ws_message_data = wsclient.DataUpdate{ SourcePath:message.Path, SourceType:message.SourceType }
                    message_data_raw,err:= ws_message_data.GetRaw()
                    if err == nil {
                        var ws_message = &wsclient.Message{DataType:"data_update",Data:message_data_raw}
                        fmt.Printf("\nStart writing\n")
                        websocket_connection.Write(ws_message)
                        fmt.Printf("\nFinish writing\n")
                        fmt.Printf("Message: %v HaveToParse: %v\n",message,message.FieldExists("HashSum"))
                    }
                  case message :=<-websocket_connection.OutChannel:
                      if message.DataType == "server_response" {
                          var response wsclient.Response
                          data:=message.Data
                          err_unmarshal:=json.Unmarshal(data, &response)
                          if err_unmarshal == nil {
                              fmt.Printf("\nMessage from server: %v\n",response)
                          }
                      } else if message.DataType == "motion" {
                          var motion activa.Motion
                          data:=message.Data
                          err_unmarshal:=json.Unmarshal(data, &motion)
                          if err_unmarshal == nil {
                              //fmt.Printf("\nNew motion %v\n", motion)
                              motions<-&motion
                          }

                      }
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
