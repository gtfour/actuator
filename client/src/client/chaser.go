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
    path1    := "/etc/passwd"
    path2    := "/etc/group"
    path3    := "/proc/net"
    messages := make(chan evebridge.CompNotes, 100)
    wp       := chase.WPCreate()

    _ = chase.Listen(path1, messages, wp)
    _ = chase.Listen(path2, messages, wp)
    _ = chase.Listen(path3, messages, wp)

    //go func(){
    //    time.Sleep( 10000 * time.Millisecond)
    //    wp.RemoveTarget("/tmp/test/test2/toremove.txt")
    //}()
    evebridge.Handle(messages)
}
