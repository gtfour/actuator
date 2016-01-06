package ws

import "fmt"
import "io"
import "log"
import "golang.org/x/net/websocket"

const channelBufSize = 100
var   maxId      int = 0

type Client struct {
    id           int
    ws           *websocket.Conn
    server       *Server
    ch           chan             *Message
    doneChannel  chan             bool
}

func NewClient (ws *websocket.Conn, server *Server) *Client {
    if ws == nil {
        panic("ws cannot be nil")
    }
    if server == nil {
        panic("server cannot be nil")
    }
    maxId++
    ch := make(chan *Message, channelBufSize)
    doneChannel := make(chan bool)
    return &Client{ maxId, ws, server, ch, doneChannel }
}

func ( c *Client ) Conn() *websocket.Conn {
    return c.ws
}

func (c *Client) Write(msg *Message) {
    select {
        case c.ch <- msg:
        default:
            c.server.Del(c)
            err := fmt.Errorf("client %d is disconnected.", c.id)
            c.server.Error(err)
        }
}

func (c *Client) Done() {
    c.doneChannel <- true
}

func (c *Client) Listen() {
    go c.listenWrite()
    c.listenRead()
}

func (c *Client) listenWrite(){
    log.Println("Listening write to client")
    for {
        select {
            case msg := <-c.ch:
                log.Println("Send:", msg)
                websocket.JSON.Send(c.ws, msg)
            case <-c.doneChannel:
                c.server.Del(c)
                c.doneChannel <- true
                return
        }
    }
}
func (c *Client) listenRead(){
    log.Println("Listening read from client")
    for {
        select {
        case <-c.doneChannel:
            c.server.Del(c)
            c.doneChannel <-true
            return
        default:
            var msg Message
            err := websocket.JSON.Receive(c.ws, msg)
            if err == io.EOF {
                c.doneChannel <- true
            } else if err != nil {
                c.server.Error(err)
            } else {
                c.server.SendAll(&msg)
            }
        }
    }
}
