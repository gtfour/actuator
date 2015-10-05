package wsserver

import (
        "log"
        "net/http"
        "code.google.com/p/go.net/websocket"
)

type Server struct {

    pattern   string
    messages  []*Message
    clients   map[int]*Client
    addCh     chan *Client
    delCh     chan *Client
    sendAllCh chan *Message
    doneCh    chan bool
    errCh chan error
}

func NewServer(pattern string) *Server {
    messages  := []*Message{}
    clients   := make(map[int]*Client)
    addCh     := make(chan *Client)
    delCh     := make(chan *Client)
    sendAllCh := make(chan *Message)
    doneCh    := make(chan bool)
    errCh     := make(chan error)

    return &Server {
               pattern,
               messages,
               clients,
               addCh,
               delCh,
               sendAllCh,
               doneCh,
               errCh,
           }
}

func ( server *Server ) Add ( client *Client ) {

    server.addCh <- client

}

func ( server *Server ) Del ( client *Client ) {

    server.delCh <- client

}

func ( server *Server ) SendAll ( message *Message ) {

    server.sendAllCh <- message

}

func ( server *Server ) Done () {

    server.doneCh <- true

}

func ( server *Server ) Err ( err error ) {

    server.errCh <- err

}

func ( server *Server ) sendPastMessages ( client *Client ) {

    for _,msg := range server.messages {
        client.Write(msg)
    }

}

func ( server *Server ) sendAll ( message *Message ) {

    for _, client := range server.clients {
        client.Write( message )
    }

}

func ( server *Server ) Listen () {

    log.Println("Listening server...")

    onConnected := func( ws *websocket.Conn ) {
        defer func() {
            err := ws.Close()
            if err!= nil {
                server.errCh <- err
            }
        }()
            client := NewClient( ws, server )
            server.Add(client)
            client.Listen()
   }

    http.Handle(server.pattern, websocket.Handler(onConnected))
    log.Println("Created handler")
    for {
        select {

        case client := <-server.addCh:

            log.Println("Added new client")
            server.clients[client.id] = client
            log.Println("Now", len(server.clients), "clients connected.")
            server.sendPastMessages(client)

        case client:= <-server.delCh:
            log.Println("Delete client")
            delete(server.clients, client.id)

        case msg := <-server.sendAllCh:
             log.Println("Send all:", msg )
             server.messages = append(server.messages, msg)
             server.sendAll(msg)

        case err := <-server.errCh:
              log.Println("Error:", err.Error())

        case <-server.doneCh:
              return
        }
    }
}



