package wsserver

import "io"
import "log"
import "fmt"
import "strconv"
//import "encoding/json"
import "golang.org/x/net/websocket"

// import "wengine/activa"
import "wengine/dusk"
// import "jumper/common/marconi"

import "jumper/common/gen"


const channelBufSize      =  100
var   maxIndexNumber int  =  0
var   database            =  dusk.DATABASE_INSTANCE

type Client struct {
    //
    //
    //
    IndexNumber  int
    ws           *websocket.Conn
    server       *Server
    ch           chan             *Message
    doneChannel  chan             bool
    session_id   string
    name         string
    id           string
    //
    //
    //
}

func NewClient ( ws *websocket.Conn, server *Server )( *Client ) {
    //
    //
    if ws == nil {
        panic("ws cannot be nil")
    }
    if server == nil {
        panic("server cannot be nil")
    }
    //
    //
    maxIndexNumber++
    ch            := make(chan *Message, channelBufSize)
    doneChannel   := make(chan bool)
    //
    // temporary solution
    //
    maxIdStr      := strconv.Itoa(maxIndexNumber)
    newClientId,_ := gen.GenId()
    //
    //
    return &Client{ maxIndexNumber, ws, server, ch, doneChannel, "", maxIdStr, newClientId } // session_id is empty yet . Will be filled when recieve first "ws_state":"open" message
    //
    //
}

func (c *Client)Write(msg *Message) {
    select {
        case c.ch <- msg:
        default:
            c.server.Del(c)
            err := fmt.Errorf("client %d is disconnected.", c.IndexNumber)
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
    //
    log.Println("\nListening read from client\n")
    //
    for {
        select {
        case <-c.doneChannel:
            c.server.Del(c)
            c.doneChannel <-true
            return
        default:
            //var msg MessageType
            var msg Message
            log.Println("Awaiting message from clients...")
            err := websocket.JSON.Receive(c.ws, &msg)  // waiting
            log.Println("Recieve has been finished. Processing...")
            //
            //  fmt.Printf("\nMessage : %v\n",msg)
            //
            if err == io.EOF {
                c.doneChannel <- true
            } else if err != nil {
                c.server.Error(err)
            } else {
                _ = c.handleMessage(&msg)
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

//
//
//
/*
func (c *Client)handleMessage(msg *Message)(err error){
    //
    //
    //
    switch data_type := msg.DataType; data_type {
        //
        // according to data_type we convert Data field to appropriate message type 
        //
        case "data_update":
            //
            var msg_du marconi.DataUpdate
            data          := msg.Data
            err_unmarshal := json.Unmarshal( data, &msg_du )
            //
            //
            //
            if err_unmarshal == nil && msg_du.SourcePath != "/tmp/test/motion.test" {
                //
                // c.server.SendAll(&msg_chat)
                //
                fmt.Printf("\n<Message Data Update: %v\n",msg_du)
                var response      Message
                var response_data marconi.Response
                response_data.Status  =  marconi.STATUS_OK
                response.DataType     =  "server_response"
                response_data_raw,err := response_data.GetRaw()
                fmt.Printf("\nStatus message len %v\n",len(response_data_raw))
                if err == nil {
                    fmt.Printf("\n<<Sending response>>\n")
                    response.Data = response_data_raw
                    c.Write(&response)
                }
                //
                //
                //
            }
            //
            //
            //
            if msg_du.SourcePath == "/tmp/test/motion.test" {
                motion := activa.CreateMotion()
                database.WriteMotion(&motion)
                var response      Message
                response.DataType     =  "motion"
                response_data_raw,err := motion.GetRaw()
                if err == nil {
                    fmt.Printf("\n:: Sending motion ::\n")
                    response.Data = response_data_raw
                    c.Write(&response)
                }
            }
            //
            //
            //
        case "message_ws_state":
            //
        case "new_dynima":
            //
    }
    return nil
}
*/
