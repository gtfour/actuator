package evebridge

import "fmt"
import "time"
import "encoding/json"
import "client/wsclient"
// // import "client/activa"
import "jumper/activa"
// // import "client/cross"
import "client/majesta"
import "jumper/common/marconi"

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
        //
        //
        //
        fmt.Printf("\n:: Start handling ::\n")
        motions := make(chan *activa.Motion, 100)
        //
        // // go activa.Handle(motions)
        //
        fmt.Printf("\nTrying to get ws-connection...\n")
        var websocket_connection = wsclient.GetWsConnection()
        fmt.Printf("\n--ws-connection open error: %v --\n", websocket_connection.OpenError)
        //
        //
        //
        for {
            select{
                case message :=<-messages:
                    //
                    //
                    fmt.Printf("\n<<Evebridge: message has been recieved>>\n")
                    //
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
                    //
                case message :=<-websocket_connection.OutChannel:
                      //
                      switch message_type := message.DataType; message_type {
                          //
                          case "server_response":
                              //
                              //
                              var response wsclient.Response
                              data          := message.Data
                              err_unmarshal := json.Unmarshal( data, &response )
                              if err_unmarshal == nil {
                                  fmt.Printf("\nMessage from server: %v\n", response)
                              }
                              //
                              //
                          case "motion":
                              //
                              // motion is instruction about which file should be modified
                              //
                              var motion activa.Motion
                              data           := message.Data
                              err_unmarshal  := json.Unmarshal(data, &motion)
                              //
                              //
                              if err_unmarshal == nil {
                                  //
                                  // fmt.Printf("\nNew motion %v\n", motion)
                                  //
                                  motion.TaskState =  activa.TASK_STATE_inprogress
                                  motionSubType    := motion.SubType
                                  motionSourceType := motion.SourceType
                                  motionSourcePath := motion.SourcePath
                                  _,_              =  motionSourceType, motionSourcePath
                                  // // cross.WriteMotion( &motion )
                                  //
                                  //
                                  switch motion_type := motion.Type; motion_type {
                                      //
                                      //
                                      case activa.MOTION_TYPE_BLACKOUT:
                                      //
                                      // // commands

                                      //
                                      //
                                      case activa.MOTION_TYPE_BLACKTOP:
                                      //
                                      // // files and directories
                                      if motionSubType == activa.MOTION_SUBTYPE_ADD_DYNIMA {
                                          fmt.Printf("\nMotion %v  Data: motionSubType: %v  motionSubType: %v motionSourceType: %v motionSourcePath: %v \n", motion_type, motionSubType, motionSourceType, motionSourcePath )
                                      }
                                      //
                                      //
                                  }
                                  //
                                  //
                                  //
                                  motions<-&motion
                              }
                          case "dynima":
                              var request marconi.Request
                              data           := message.Data
                              err_unmarshal  := json.Unmarshal( data, &request )
                              //
                              // way to get here:
                              //    send post request to wengine server:
                              //    curl --data 'dashboardName="Users";sourceType="TARGET_FILE";sourcePath="/etc/passwd";clientName=1' "http://127.0.0.1:9000/rest/dashboard/add-dashboard/"
                              //    
                              // seems now we can recieve params on client side by sending json-requests to wengine master server :))))
                              // example of curl-request is above
                              // clientName is index number assigned to websocket client on first connection and equals to maxIndexNumber global variable
                              //
                              //
                              if err_unmarshal == nil {
                                  fmt.Printf("\nHandling dynima:\n")
                                  fmt.Printf("params:\nChangeType: %d\nObjName: %s\nObjType: %s\nObjPath: %s\n", request.ChangeType,request.ObjName, request.ObjType, request.ObjPath)
                              }
                              //
                              //
                      }
                default:
                    time.Sleep( LOG_CHANNEL_TIMEOUT_MS  * time.Millisecond )
                    fmt.Println("No messages")
            }
        }
}
