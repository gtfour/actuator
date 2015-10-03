package wsclient

import ( "fmt"
         "io"
         "log"
         "code.google.com/p/go.net/websocket"
       )

const channelBufsize = 100

var maxId int = 0

type Client struct {

    id      int
    ws      *websocker.Conn
    server  *Server
    ch      chan *Message
    doneCh  chan bool

}

func NewClient ( ws *websocket.Conn, server *Server ) *Client {

    if ws     == nil { panic("ws cannot be nil"    ) }
    if server == nil { panic("server cannot be nil") }

    maxId ++

    ch     := make( chan *Message, channelBufsize )
    doneCh := make( chan bool )

    return &Client{ maxId, ws, server, ch, doneCh }
}
