package evebridge

import "client/wsclient"
import "jumper/common/marconi"
import "jumper/cuda/targets"


func (a *App)handleDynima(request *marconi.Request)(){
    FUNCTION_NAME:="handleDynima"
    var newTargetType string = targets.TARGET_UNDEFINED_STR
    //
    if request.ObjType == targets.TARGET_FILE_STR {
        //
        newTargetType = targets.TARGET_FILE_STR
        //
    } else if request.ObjType == targets.TARGET_DIR_STR {
        //
        newTargetType = targets.TARGET_DIR_STR
        //
    } else {
        err:=targetTypeUndefined
        a.writeLogEntry(PACKAGE_NAME,FUNCTION_NAME,"obj type check",err)
        return
    }
    targetConfig          :=  make(map[string]string, 0)
    targetConfig["type"]  =   newTargetType
    targetConfig["path"]  =   request.ObjPath
    tgt,err               :=  targets.NewTarget(targetConfig)
    _ = tgt
    if err != nil {
        a.writeLogEntry(PACKAGE_NAME,FUNCTION_NAME,"new target config",err)
        return
    }
    //
    response_type        := "dynima_response"
    _ = response_type
    var message_data_raw []byte
    var ws_message = &wsclient.Message{DataType:"data_update",Data:message_data_raw}
    a.websocketConn.Write(ws_message)

}
