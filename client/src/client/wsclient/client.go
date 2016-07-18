//package main
package wsclient

import "fmt"
import "log"
import "encoding/json"
import "golang.org/x/net/websocket"
import "client/settings"

var WsConn = CreateConnection(settings.RESTAPI_WS_ORIGIN, settings.RESTAPI_WS_URL)


type Event struct {

    Author string `json:"author"`
    Body   string `json:"body"`
    //Name     string
    //Type     string
    //FilePath string
    //Hostname string
    //Os       string
    //Version  string
    //Release  string

}

type WebSocketConnection struct {

    ws         *websocket.Conn
    InChannel  chan *Message
    OutChannel chan *Message
    OpenError  error

}

func (wsconn *WebSocketConnection) Handle()(error) {
    if wsconn.OpenError == nil {
        go wsconn.Read()
        for {
            select {
                case message :=<-wsconn.InChannel:
                    message_raw,_:=message.GetRaw()
                    if test,err := wsconn.ws.Write(message_raw) ; err != nil {
                        fmt.Println(test)
                        log.Fatal(err)
                    }
                /*case message :=<-wsconn.OutChannel:
                    var response Response
                    data:=message.Data
                    err_unmarshal:=json.Unmarshal(data, &response)
                    if err_unmarshal == nil {
                        fmt.Printf("\nMessage from server: %v\n",response)
                    }
                */
            }
        }
    } else {
        fmt.Printf("\nError while opening ws connection\n")

    }
    return nil
}

func (wsconn *WebSocketConnection) Write(m *Message )( ) {
    wsconn.InChannel <- m
}

func (wsconn *WebSocketConnection) Read()(error) {
    for {
        var data []byte
        data = make([]byte, 512) // 53
        var message Message
        n, err := wsconn.ws.Read(data) // replace n to _
        fmt.Printf("\n<<Reading1>>\nErr:%v\ndata:%v\nread data:%v\n",err,data,n)
        if err!= nil { return err }
        err    = json.Unmarshal(data[:n], &message)
        fmt.Printf("\n<<Reading2>>\nErr:%v\n",err)
        if err == nil { wsconn.OutChannel <- &message  }
    }
}




func CreateConnection ( origin string, url string ) (  *WebSocketConnection ) {

    fmt.Printf("Creating connection:\norigin:%v\nurl:%v\n",origin,url)
    wsConn    :=  &WebSocketConnection{}
    protocol  :=  ""
    ws,err    :=  websocket.Dial( url, protocol, url )

    if err != nil {
       wsConn.OpenError = err
       return wsConn

    } else {
        wsConn.ws         = ws
        wsConn.InChannel  = make(chan *Message, 100)
        wsConn.OutChannel = make(chan *Message, 100)
        wsConn.OpenError  = nil
        go wsConn.Handle()
        fmt.Printf("\nws-connection has been opened\n")
        return wsConn
    }
}




/*func main() {

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

}*/
