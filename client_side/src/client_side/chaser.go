package main

import _ "net/http/pprof"
import "net/http"
import "fmt"
//
import "client_side/chase"
import "time"

func main() {

    path     := "/proc/net"
    messages :=   make(chan string,100)
    wp       :=  chase.WPCreate()
    _ = chase.Listen(path, messages, wp)

    go func() {
        fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
    }()

    for {

        select{
            case message:=<-messages:
                fmt.Println(message)

            default:
                time.Sleep( chase.LOG_CHANNEL_TIMEOUT_MS * time.Millisecond)
                fmt.Println("No messages")

        }

    }

}
