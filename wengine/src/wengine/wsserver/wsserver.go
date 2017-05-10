package wsserver

import "fmt"
import "log"
import "golang.org/x/net/websocket"
import "github.com/gin-gonic/gin"
import "wengine/settings"
import "wengine/rest"
// import "golang.org/x/net/websocket"

var WebSocketServerWeb = CreateServer( settings.WS_WEBDATA_URL )
var WebSocketServerSrv = CreateServer( settings.WS_SRVDATA_URL )

type Server struct {

    pattern        string
    messages       []*Message
    Clients        map[int]*Client
    addChannel     chan *Client
    delChannel     chan *Client
    sendAllChannel chan *Message
    doneChannel    chan bool
    errChannel     chan error
    WShandler      websocket.Handler

}

func(s *Server)GetHandler(c *gin.Context)(function func(ws *websocket.Conn)){

    token,user,err:=rest.GetTokenFromCookies(c)
    fmt.Printf("token:%v\nuser:%v\nerr:%v\n",token,user,err)

    function = func(ws *websocket.Conn){
        defer func(){
            err := ws.Close()
            if err != nil {
                s.errChannel <- err
            }
        }()
        client := NewClient(ws, s)
        s.Add(client)
        client.Listen()
    }

    return function
}


func CreateServer (pattern string) *Server {
    messages       := []*Message{}
    clients        := make(map[int]*Client)
    addChannel     := make(chan *Client)
    delChannel     := make(chan *Client)
    sendAllChannel := make(chan *Message)
    doneChannel    := make(chan bool)
    errChannel     := make(chan error)
    // gen ws handler
    s:=&Server{
        pattern:pattern,
        messages:messages,
        Clients:clients,
        addChannel:addChannel,
        delChannel:delChannel,
        sendAllChannel:sendAllChannel,
        doneChannel:doneChannel,
        errChannel:errChannel,
        //s.WShandler:handler,
    }

    //onConnected := func(ws *websocket.Conn){
    //    defer func(){
    //        err := ws.Close()
    //        if err != nil {
    //            s.errChannel <- err
    //        }
    //    }()
    //    client := NewClient(ws, s)
    //    s.Add(client)
    //    client.Listen()
   // }
    //s.WShandler = websocket.Handler(onConnected)
    // test
    go s.Listen()
    log.Printf("Server has been created :)\nServer is listening connection :)\n")
    // test
    return s
}
//
//
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

func (s *Server) SendToAllClients (msg *Message) {
    for _, c := range s.Clients {
        c.Write(msg)
    }
}
//
//
//
func (s *Server) GetClients ()( map[int]*Client ) {
     return s.Clients
}
//
//
//
func (s *Server) GetClientByName ( client_name string )( *Client ) {
    //
    for _,c := range s.Clients {
        name := c.name
        if name == client_name {
            return c
        }
    }
    return nil
    //
}

//
//

func (s *Server)Listen(){
    //
    log.Println( "... server is listening ..." )
    //
    //
    for {
        //
        //
        //
        select {
            case c:= <-s.addChannel:
                log.Println("Added new client")
                s.Clients[c.Id] = c
                log.Println("Now", len(s.Clients), "clients connected")
            case c:= <-s.delChannel:
                log.Printf("Delete client %s",c.Id)
                delete(s.Clients, c.Id)
            case msg:= <-s.sendAllChannel:
                log.Println("Send all:", msg)
                s.messages = append(s.messages, msg)
                s.SendToAllClients(msg)
            case err:= <-s.errChannel:
                log.Println("Error:", err.Error())
            case <-s.doneChannel:
                return
        }
        //
        //
        //
    }
    //
    //
}

func WebSocketHandle(data gin.H)(func(c *gin.Context)){
    return func(c *gin.Context){
        handler   := websocket.Handler(WebSocketServerWeb.GetHandler(c))
        handler.ServeHTTP(c.Writer, c.Request)
    }
}
