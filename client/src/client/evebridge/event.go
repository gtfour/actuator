package evebridge

import "fmt"
import "time"
import "encoding/json"
import "client/wsclient"
import "client/activa"
import "client/cross"
import "client/majesta"

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


func Handle(messages chan majesta.CompNotes )() {
        fmt.Printf("\n:: Start handling ::\n")
        motions := make(chan *activa.Motion, 100)
        go activa.Handle(motions)
        fmt.Printf("\nTrying to get ws-connection...\n")
        var websocket_connection = wsclient.GetWsConnection()
        fmt.Printf("\n--ws-connection open error: %v --\n",websocket_connection.OpenError)
        for {
            select{
                case message :=<-messages:
                    fmt.Printf("\n<<Evebridge: message has been recieved>>\n")
                    var ws_message_data  =  wsclient.DataUpdate{ SourcePath:message.Path, SourceType:message.SourceType }
                    message_data_raw,err := ws_message_data.GetRaw()
                    //
                    if err == nil {
                        var ws_message = &wsclient.Message{DataType:"data_update",Data:message_data_raw}
                        fmt.Printf("\nStart writing\n")
                        websocket_connection.Write(ws_message)
                        fmt.Printf("\nFinish writing\n")
                        fmt.Printf("Message: %v HaveToParse: %v\n",message,message.FieldExists("HashSum"))
                    }
                    //
                case message :=<-websocket_connection.OutChannel:
                      //
                      switch message_type := message.DataType; message_type {
                          case "server_response":
                              var response wsclient.Response
                              data          := message.Data
                              err_unmarshal := json.Unmarshal(data, &response)
                              if err_unmarshal == nil {
                                  fmt.Printf("\nMessage from server: %v\n",response)
                              }
                          case "motion":
                              var motion activa.Motion
                              data          := message.Data
                              err_unmarshal := json.Unmarshal(data, &motion)
                              if err_unmarshal == nil {
                                  //fmt.Printf("\nNew motion %v\n", motion)
                                  motion.TaskState=activa.TASK_STATE_inprogress
                                  cross.WriteMotion(&motion)
                                  //
                                  if motion.Type == activa.MOTION_TYPE_BLACKOUT {

                                  } else if motion.Type == activa.MOTION_TYPE_BLACKTOP {

                                  }
                                  //
                                  motions<-&motion
                              }


                      }
                default:
                    time.Sleep( LOG_CHANNEL_TIMEOUT_MS  * time.Millisecond )
                    fmt.Println("No messages")
            }
        }
}
