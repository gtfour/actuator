package main

import _ "net/http/pprof"
import "net/http"
import "fmt"
//
import "client/chase"
import "client/evebridge"
// import "time"

func main() {

    path     := "/tmp/test"
    messages :=   make(chan evebridge.CompNotes,100)
    wp       :=  chase.WPCreate()
    _ = chase.Listen(path, messages, wp)

    go func() {
        fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
    }()
    evebridge.Handle(messages)


}