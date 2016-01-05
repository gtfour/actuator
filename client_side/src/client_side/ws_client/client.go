package main
//package wsclient

import (
       "code.google.com/p/go.net/websocket"
       "fmt"
       "log"
       "encoding/json"
)

type Event struct {

    Author string `json:"author"`
    Body string `json:"body"`
    //Name     string
    //Type     string
    //FilePath string
    //Hostname string
    //Os       string
    //Version  string
    //Release  string

}

func CreateConnection ( origin, url string ) ( ws *websocket.Conn, err error ) {

    protocol  :=  ""

    ws,err    =  websocket.Dial( url, protocol, url )

    if err != nil {

       return

    }

    return
}




func main() {

    origin    :=  "http://127.0.0.1"
    url       :=  "ws://127.0.0.1:8090/entry"

    ws,err :=  websocket.Dial( url, "", origin )

    if err != nil {

        log.Fatal(err)

    }

    test:=&Event{Author:"venom",Body:"zenity-3.4.0-0ubuntu4"}

    test_serialized, err := json.Marshal(test)

    fmt.Println( err )
    fmt.Println( string(test_serialized) )

    if test,err := ws.Write( test_serialized ) ; err != nil {

        fmt.Println(test)
        log.Fatal(err)

    }

    var msg = make( []byte, 512 )
    var n int

    if n, err = ws.Read(msg) ; err!= nil {

        log.Fatal(err)

    }

    fmt.Printf("Received: %s.\n",msg[:n])

}
