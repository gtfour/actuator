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
    messages  := []*Messages{}
    clients   := make(map[int])
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

func ( s *Server ) Add ( client *Client ) {

    server.addCh <- client

}
