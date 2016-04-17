package ws

import "fmt"
import "io"
import "log"
import "encoding/json"
import "golang.org/x/net/websocket"

const channelBufSize = 100
var   maxId      int = 0

type Client struct {
    id           int
    ws           *websocket.Conn
    server       *Server
    ch           chan             *MessageChat
    doneChannel  chan             bool
    session_id   string
}

func NewClient (ws *websocket.Conn, server *Server) *Client {
    if ws == nil {
        panic("ws cannot be nil")
    }
    if server == nil {
        panic("server cannot be nil")
    }
    maxId++
    ch := make(chan *MessageChat, channelBufSize)
    doneChannel := make(chan bool)
    return &Client{ maxId, ws, server, ch, doneChannel }
}

func ( c *Client ) Conn() *websocket.Conn {
    return c.ws
}

func (c *Client) Write(msg *MessageChat) {
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
            //var msg MessageType
            var msg Message
            err := websocket.JSON.Receive(c.ws, &msg)
            fmt.Printf("\nMessage : %v\n",msg)

            if err == io.EOF {
                c.doneChannel <- true
            } else if err != nil {
                c.server.Error(err)
            } else {
                data_type := msg.DataType
                if data_type == "message_chat" {
                    var msg_chat MessageChat
                    data:=msg.Data
                    err_unmarshal:=json.Unmarshal(data, &msg_chat)
                    if err_unmarshal!= nil {
                        c.server.SendAll(&msg_chat)
                    }
                } else if data_type == "message_switch_dashboard" {
                    var msg_swd MessageSwitchDashboard
                    data:=msg.Data
                    //source := (*json.RawMessage)(&data)
                    //data_byte:=[]byte {}
                    //err_unmarshal:=json.Unmarshal(data, &msg_swd)
                    //err_unmarshal_raw:=data.UnmarshalJSON(data_byte)
                    err_unmarshal:=json.Unmarshal(data, &msg_swd)
                    if err_unmarshal == nil {
                        //c.server.SendAll(&msg_chat)
                        fmt.Printf("\nMessage Switch Dashboard: %v\n",msg_swd)
                    }

                }
                //c.server.SendAll(&msg)
            }
        }
    }
}
