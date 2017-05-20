package main

//import _ "net/http/pprof"
//import "net/http"
import "fmt"
import "client/chase"
import "client/majesta"
import "client/evebridge"
//import "time"

// access to /tmp/cross.db may hang this app

func main() {
    fmt.Printf("\n::main has been started\n")
    //
    //go func() {
    //     fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
    //}()
    path1      := "/tmp/test"
    //path2    := "/etc/group"
    //path3    := "/proc/net"
    messages   := make(chan majesta.CompNotes, 100)
    fmt.Printf("\n:: creating worker-pool:: \n")
    wp       := chase.WPCreate()
    fmt.Printf("\n:: worker pool has been created:: \n")

    fmt.Printf("\n::start chasing ..\n")
    _ = chase.Listen(path1, messages, wp)
    fmt.Printf("\n::chase has been started ..\n")
    //_ = chase.Listen(path2, messages, wp)
    //_ = chase.Listen(path3, messages, wp)
    //go func(){
    //    time.Sleep( 10000 * time.Millisecond)
    //    wp.RemoveTarget("/tmp/test/test2/toremove.txt")
    //}()
    evebridge.Handle(messages)
}
