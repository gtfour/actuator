package evebridge

import "fmt"
import "client/wsclient"
import "jumper/common/marconi"
import "jumper/cuda"
import "jumper/cuda/targets"
import "jumper/cuda/filtering"


func (a *App)handleDynima(request *marconi.Request)(){
    var err error
    FUNCTION_NAME := "handleDynima"
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
        err=targetTypeUndefined
        //a.writeLogEntry(PACKAGE_NAME,FUNCTION_NAME,"obj type check",err)
        return
    }
    a.writeLogEntry(PACKAGE_NAME,FUNCTION_NAME,"obj type check",err)
    targetConfig          :=  make(map[string]string, 0)
    targetConfig["type"]  =   newTargetType
    targetConfig["path"]  =   request.ObjPath
    tgt,err               :=  targets.NewTarget(targetConfig)
    _ = tgt
    //
    if err != nil {
        a.writeLogEntry(PACKAGE_NAME,FUNCTION_NAME,"new target config",err)
        return
    }
    //
    // temporary code block
    //
    d                  :=  cuda.Dynima{}
    defaultFilterList  :=  filtering.CreateDefaultFilterList()
    for i := range filtering.CreateDefaultFilterList() {
        filter := defaultFilterList[i]
        d.AppendFilter( filter )
    }
    d.AppendTarget(tgt)
    resultSet       := d.RunFilters()
    //
    // checking result before sending to wengine
    //
    // //results,err:=resultSet.GetData()
    // //fmt.Printf("\n==checking result set==\n")
    // //for i:= range results {
    // //    r:=results[i]
    // //    fmt.Printf("\n%v",r)
    // //}
    // //fmt.Printf("\n== ==\n")
    //
    //
    //
    result_byte,err := resultSet.GetJson()
    if err != nil {
        a.writeLogEntry(PACKAGE_NAME,FUNCTION_NAME,"converting dynima result to byte",err)
        return
    }
    //
    fmt.Printf("\n<< -- result_byte: %v -- >>\n", result_byte)
    //
    //
    response_type        := "dynima_response"
    var ws_message       =  &wsclient.Message{DataType:response_type, Data:result_byte}
    //fmt.Printf("\n-- Sending response back to server --\n")
    a.writeLogEntry(PACKAGE_NAME,FUNCTION_NAME,"sending response back to server",nil)
    a.websocketConn.Write(ws_message)
    //
    //
}
