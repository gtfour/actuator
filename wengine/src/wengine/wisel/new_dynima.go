package wisel

import "wengine/wsserver"
import "wengine/core/marconi"

func AddNewDynima(params map[string]string)(error) {
    //
    //
    dashboardName,okDashboardName  :=  params["dashboardName"]
    sourceType,okSourceType        :=  params["sourceType"]
    sourcePath,okSourcePath        :=  params["sourcePath"]
    clientName,okClientName        :=  params["clientName"]
    clientId,okClientId            :=  params["clientId"]
    //
    //
    if okDashboardName && okSourceType && okSourcePath && ( okClientName || okClientId ) {
        //
        var request  marconi.Request
        var wsClient *wsserver.Client
        var wsMessage wsserver.Message
        //
        wsServer := wsserver.WebSocketServerWeb
        //
        if okClientId {
            wsClient = wsServer.GetClientById(clientId)
        } else if okClientName {
            wsClient = wsServer.GetClientByName(clientName)
        } else {
            wsClient = nil
            return PARAMS_NOT_ENOUGH
        }
        if wsClient == nil {  return CANT_FIND_WSCLIENT  }
        //
        request.ChangeType =  marconi.CTYPE_ADD_NEW_DYNIMA
        request.ObjName    =  dashboardName
        request.ObjType    =  sourceType
        request.ObjPath    =  sourcePath
        request_byte,err   := request.GetRaw()
        //
        if err == nil {
            //
            wsMessage.DataType = "dynima"
            wsMessage.Data     = request_byte
            wsClient.Write( &wsMessage )
            return nil
            //
        } else {
            //
            return err
            //
        }
        //
    } else {
        //
        return PARAMS_NOT_ENOUGH
        //
    }
    //
}
