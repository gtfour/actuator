package main
//package wsclient

import (
       "code.google.com/p/go.net/websocket"
       "fmt"
       "log"
)

func main() {

    url    := "http://127.0.0.1"
    origin := "ws://127.0.0.1:8080/entry"

    ws,err := websocket.Dial(origin, "", url)

    if err != nil {

        log.Fatal(err)

    }

    if test,err := ws.Write([]byte(`{"author":"venom","body":"Buy cheese and bread for breakfast."}`)) ; err != nil {

        fmt.Println(test)
        log.Fatal(err)
    }

    var msg = make([]byte, 512)
    var n int

    if n, err = ws.Read(msg) ; err!= nil {
        log.Fatal(err)
    }

    fmt.Printf("Received: %s.\n",msg[:n])

}
