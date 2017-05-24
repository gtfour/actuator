package evebridge

import "client/wsclient"
import "client/majesta"


func (a *App)handleTextOutput(message *majesta.CompNotes)(){
    FUNCTION_NAME:="handleTextOutput"
    // fmt.Printf("\n<<Evebridge: message has been recieved>>\n")
    var ws_message_data  =  wsclient.DataUpdate{ SourcePath:message.Path, SourceType:message.SourceType }
    message_data_raw,err := ws_message_data.GetRaw()
    if err == nil {
        var ws_message = &wsclient.Message{DataType:"data_update",Data:message_data_raw}
        // // fmt.Printf("\nStart writing\n")
        a.writeLogEntry(PACKAGE_NAME, FUNCTION_NAME, "start writing message to websocket",nil)
        a.websocketConn.Write(ws_message)
        a.writeLogEntry(PACKAGE_NAME, FUNCTION_NAME, "writing has been finished",nil)
        // // fmt.Printf("\nFinish writing\n")
        // // fmt.Printf("Message: %v HaveToParse: %v\n",message,message.FieldExists("HashSum"))
    } else {
        a.writeLogEntry(PACKAGE_NAME, FUNCTION_NAME,"convert message to ws-message",err)
    }
    return
}
