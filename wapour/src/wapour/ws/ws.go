package ws

import "log"
//import "net/http"
import "fmt"
import "golang.org/x/net/websocket"
import "github.com/gin-gonic/gin"

type Server struct {

    pattern        string
    messages       []*Message
    clients        map[int]*Client
    addChannel     chan *Client
    delChannel     chan *Client
    sendAllChannel chan *Message
    doneChannel    chan bool
    errChannel     chan error
    WShandler      websocket.Handler

}

func NewServer (pattern string) *Server {
    messages       := []*Message{}
    clients        := make(map[int]*Client)
    addChannel     := make(chan *Client)
    delChannel     := make(chan *Client)
    sendAllChannel := make(chan *Message)
    doneChannel    := make(chan bool)
    errChannel     := make(chan error)
    // gen ws handler
    s:=&Server {
        pattern:pattern,
        messages:messages,
        clients:clients,
        addChannel:addChannel,
        delChannel:delChannel,
        sendAllChannel:sendAllChannel,
        doneChannel:doneChannel,
        errChannel:errChannel,
        //s.WShandler:handler,
    }

    onConnected := func(ws *websocket.Conn) {
        defer func() {
            err := ws.Close()
            if err != nil {
                s.errChannel <- err
            }
        }()
        client := NewClient(ws, s)
        s.Add(client)
        client.Listen()
    }
    //http.Handle(s.pattern, websocket.Handler(onConnected))
    s.WShandler = websocket.Handler(onConnected)
    //s.WShandler := handler

    return s
}

func (s *Server) Add (c *Client) {
    s.addChannel <- c
}

func (s *Server) Del (c *Client) {
    s.delChannel <- c
}
func (s *Server) SendAll(msg *Message) {
    s.sendAllChannel <- msg
}
func (s *Server) Done () {
    s.doneChannel <- true
}
func (s *Server) Error (err error) {
    s.errChannel <- err
}
func (s *Server) sendPastMessages (c *Client) {
    for _,msg := range s.messages {
        c.Write(msg)
    }
}

func (s *Server) SendToAllClients (msg *Message) {
    for _, c := range s.clients {
        c.Write(msg)
    }
}

func (s *Server) Listen() {
    log.Println("Listening server...")
    //onConnected := func(ws *websocket.Conn) {
    //    defer func() {
    //        err := ws.Close()
    //        if err != nil {
    //            s.errChannel <- err
    //        }
    //    }()
    //    client := NewClient(ws, s)
    //    s.Add(client)
    //    client.Listen()
    //}
    //http.Handle(s.pattern, websocket.Handler(onConnected))
    //s.WShandler = websocket.Handler(onConnected)
    log.Println("Created handler")
    for {
        select {
            case c:= <-s.addChannel:
                log.Println("Added new client")
                s.clients[c.id] = c
                log.Println("Now", len(s.clients), "clients connected")
                s.sendPastMessages(c)
            case c:= <-s.delChannel:
                log.Println("Delete client")
                delete(s.clients, c.id)
            case msg := <-s.sendAllChannel:
                log.Println("Send all:", msg)
                s.messages = append(s.messages, msg)
                s.SendToAllClients(msg)
            case err:= <-s.errChannel:
                log.Println("Error:", err.Error())
            case <-s.doneChannel:
                return
        }
    }
}


func WSserver(data gin.H, wshandler websocket.Handler)( func(c *gin.Context) ) {

    //server := NewServer()
    return func(c *gin.Context)  {
        //wshandler(c.Writer, c.Request)
        //serveWs(c.Writer, c.Request)
        fmt.Printf("\n<ws handler is working>\n")
        handler := wshandler
        handler.ServeHTTP(c.Writer, c.Request)
    }
}
