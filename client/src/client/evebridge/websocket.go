package evebridge

import "fmt"
import "encoding/json"
import "client/wsclient"
import "jumper/activa"
import "jumper/common/marconi"
// import "jumper/cuda/targets"
// import "client/activa"
// // import "client/majesta"
// // import "client/logging"


func (a *App)handleWebSocketConnection(message *wsclient.Message)(){
    //
    FUNCTION_NAME := "handleWebSocketConnection"
    //
    switch message_type := message.DataType; message_type {
        //
        //
        //
        case "server_response":
            //
            var response wsclient.Response
            data          := message.Data
            err_unmarshal := json.Unmarshal( data, &response )
            if err_unmarshal == nil {
                //
                fmt.Printf("\nMessage from server: %v\n", response)
                //
            } else {
                //
            }
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
            //
            if err_unmarshal == nil {
                a.handleMotion(&motion)
            } else {
                a.writeLogEntry(PACKAGE_NAME,FUNCTION_NAME,"motion",err_unmarshal)
            }
        case "dynima":
            //
            //
            //
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
                // fmt.Printf("\nHandling dynima:\n")
                // fmt.Printf("params:\nChangeType: %d\nObjName: %s\nObjType: %s\nObjPath: %s\n", request.ChangeType,request.ObjName, request.ObjType, request.ObjPath)
                // var ws_message_data  =  wsclient.DataUpdate{ SourcePath:message.Path, SourceType:message.SourceType }
                // message_data_raw,err := ws_message_data.GetRaw()
                //
                // trying to get file content via dynima 
                //
                a.handleDynima(&request)
            } else {
                a.writeLogEntry( PACKAGE_NAME,FUNCTION_NAME,"dynima",err_unmarshal )
            }
            //
            //
            //
    }
    // // return nil
    //
}
