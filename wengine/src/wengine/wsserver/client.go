package wsserver

import "io"
import "log"
import "fmt"
import "encoding/json"
import "golang.org/x/net/websocket"

import "wengine/core/marconi"


const channelBufSize = 100
var   maxId      int = 0

type Client struct {
    id           int
    ws           *websocket.Conn
    server       *Server
    ch           chan             *Message
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
    ch := make(chan *Message, channelBufSize)
    doneChannel := make(chan bool)
    return &Client{ maxId, ws, server, ch, doneChannel,""} // session_id is empty yet . Will be filled when recieve first "ws_state":"open" message
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
            //var msg MessageType
            var msg Message
            err := websocket.JSON.Receive(c.ws, &msg)
            //fmt.Printf("\nMessage : %v\n",msg)

            if err == io.EOF {
                c.doneChannel <- true
            } else if err != nil {
                c.server.Error(err)
            } else {
                data_type := msg.DataType
                if data_type == "data_update" {
                    var msg_du marconi.DataUpdate
                    data:=msg.Data
                    err_unmarshal:=json.Unmarshal(data, &msg_du)
                    if err_unmarshal == nil {
                        //c.server.SendAll(&msg_chat)
                        fmt.Printf("\n<Message Data Update: %v\n",msg_du)
                        var response      Message
                        var response_data marconi.Response
                        response_data.Status = marconi.STATUS_OK
                        response.DataType    = "server_response"
                        response_data_raw,err:=response_data.GetRaw()
                        fmt.Printf("\nStatus message len %v\n",len(response_data_raw))
                        if err == nil {
                            fmt.Printf("\n<<Sending response>>\n")
                            response.Data = response_data_raw
                            c.Write(&response)
                        }
                    }

                }/* else if data_type == "message_ws_state" {
                    var msg_wsst MessageWsState
                    data:=msg.Data
                    err_unmarshal:=json.Unmarshal(data, &msg_wsst)
                    if err_unmarshal == nil {
                        //c.server.SendAll(&msg_chat)
                        fmt.Printf("\nMessage Ws State: %v\n", msg_wsst)
                        if  msg_wsst.State == "open" {
                            fmt.Printf("\n>>Adding session_id to client<<\n")
                        }
                    }
                }*/
                //c.server.SendAll(&msg)
            }
        }
    }
}
