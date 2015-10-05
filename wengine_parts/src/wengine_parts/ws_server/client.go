package wsserver

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
    messageChannel      chan *Message
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

func ( client *Client ) Conn() *websocket.Conn {
    return client.ws
}

func ( client *Client ) Write ( msg *Message ) {

    select {
        case client.messageChannel <- msg.Message:
        default:
            client.server.Del( client )
            fmt.Errorf("Client %d is disconnected", client.id)
            client.server.Err(err)
    }

}

func ( client *Client ) Done () {

    client.doneCh <- true
}

func ( client *Client ) Listen() {

    go client.listenWrite()
    client.listenRead()

}

func ( client *Client ) listenWrite () {

    log.Println(" Listening write to client ")

    for {
        select {

        case msg:= <-client.messageChannel:
            log.Println("Send:",msg)
            websocket.JSON.Send( client.ws, msg )
        case <-doneCh:
            client.server.Del(client)
            doneCh <- true
            return
        }
    }
}

func ( client *Client ) listenRead () {

    log.Println(" Listening read to client")
    for {
        select {
        case <- client.doneCh:
            client.server.Del(client)
            client.doneCh <- true
            return
        default:
            var message Message
            err := websocket.JSON.Receive(client.ws, &message)
            if err == io.EOF {
                c.doneCh <- true
            } else if err != nil {
                client.server.Err(err)
            } else {
                client.server.SendAll(&message)
            }
        }
    }
}
