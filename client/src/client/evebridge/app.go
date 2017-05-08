package evebridge

import "client/wsclient"
import "client/activa"
import "client/majesta"

type App struct {
    //
    //
    WebsocketConn   wsclient.WebSocketConnection
    FileUpdates     chan majesta.CompNotes
    CommandUpdates  chan majesta.CompNotes
    Motions         chan *activa.Motion
    //
    //
}

func (a *App)Handle()(error){
    //
    return nil
    //
}

func (a *App)handleWebSocketConnection()(error){
    //
    return nil
    //
}

func (a *App)handleFileUpdates()(error){
    //
    return nil
    //
}

func (a *App)handleCommandUpdates()(error){
    //
    return nil
    //
}

func (a *App)handleMotions()(error){
    //
    return nil
    //
}
