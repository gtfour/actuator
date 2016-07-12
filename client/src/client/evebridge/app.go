package evebridge


import "client/wsclient"
import "client/activa"

type App struct {

    WebsocketConn   wsclient.WebSocketConnection
    FileUpdates     chan CompNotes
    CommandUpdates  chan CompNotes
    Motions         chan *activa.Motion



}

func (a *App) Handle ()(error) {

    return nil


}
