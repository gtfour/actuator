package main

//import _ "net/http/pprof"
//import "net/http"
//import "fmt"
//
import "client/chase"
import "client/evebridge"
//import "time"

func main() {

    //go func() {
    //    fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
    //}()
    path     := "/proc/net"
    messages :=   make(chan evebridge.CompNotes,100)
    wp       :=  chase.WPCreate()
    _ = chase.Listen(path, messages, wp)
    //go func(){
    //    time.Sleep( 10000 * time.Millisecond)
    //    wp.RemoveTarget("/tmp/test/test2/toremove.txt")
    //}()
    evebridge.Handle(messages)
}
