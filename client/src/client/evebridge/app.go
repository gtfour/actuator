package evebridge

import "fmt"
import "time"

import "client/wsclient"
//import "client/activa"
import "jumper/activa"
import "client/majesta"
import "client/logging"

type MessageSource interface{
    //
    Orient(messages chan majesta.CompNotes)()
    //
}

type App struct {
    //
    websocketConn   *wsclient.WebSocketConnection
    fileUpdates     chan majesta.CompNotes
    //  // commandUpdates  chan majesta.CompNotes
    // majestaCommon   chan majesta.CompNotes
    motions         chan *activa.Motion
    logs            chan logging.LogMessage
    configured      bool
    //
}

func (a *App)Handle()(error){
    //
    // var websocket_connection = wsclient.GetWsConnection()
    //
    for {
        select {
            case messageF    :=<-a.fileUpdates:
                a.handleTextOutput( &messageF )
            // case messageC    :=<-a.commandUpdates:
            //     a.handleTextOutput( &messageC )

            case motion     :=<-a.motions:
                a.handleMotion( motion )

            case wsmessage  :=<-a.websocketConn.OutChannel:
                a.handleWebSocketConnection( wsmessage )

            case logmessage :=<-a.logs:
                fmt.Printf("\nLOG: %v", logmessage)

            default:
                time.Sleep( LOG_CHANNEL_TIMEOUT_MS  * time.Millisecond )

        }
    }
    return nil
    //
}

/*
func (a *App)handleMotions()(error){
    //
    return nil
    //
}
func (a *App)handleLogs()(error){
    //
    return nil
    //
}
*/

func (a *App)writeLogEntry( packageName,functionName,procedureName string, err error )( error ){
    //
    message := logging.LogMessage{PackageName:packageName, FunctionName:functionName, ProcedureName:procedureName, Err:err}
    a.logs  <- message
    //
    return nil
    //
}


func MakeApp()( *App, error ){
    //
    var app App
    //
    app.fileUpdates     = make( chan majesta.CompNotes,  100 )
    //
    // app.commandUpdates  = make( chan majesta.CompNotes,  100 )
    //
    app.motions         = make( chan *activa.Motion,     100 )
    app.logs            = make( chan logging.LogMessage, 100 )
    //
    app.websocketConn   = wsclient.GetWsConnection()
    app.configured      = true
    //
    return &app, nil
    //
    //
}

func (a *App)AddMessageSource( m MessageSource )( error ){
    if a.configured {
        m.Orient(a.fileUpdates)
        return nil
    } else {
        return appWasNotConfigured
    }
}
