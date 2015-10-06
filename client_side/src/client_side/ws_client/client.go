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
    url       :=  "ws://127.0.0.1:8080/entry"

    ws,err :=  websocket.Dial( url, "", origin )

    if err != nil {

        log.Fatal(err)

    }

    test:=&Event{Author:"venom",Body:"Package json implements encoding and decoding of JSON objects as defined in RFC 4627.\n The mapping between JSON objects and Go values is described in the documentation for the Marshal and Unmarshal functions.\nParsing a format like JSON in a statically typed language like Go presents a bit of a problem.\nIf anything could show up in the JSON body, how does the compiler know how to setup memory to have a spot to place everything?\n"}

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
