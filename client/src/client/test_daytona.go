package main
//
//
//
import "fmt"
import "client/chase"
import "client/majesta"
import "client/evebridge"
//
//
// mixing chasing ( by chase )  and parsing ( by cuda )
//
//
func main() {
    //
    fmt.Printf("\n:: main has been started\n")
    path1      := "/tmp/test"
    messages   := make(chan majesta.CompNotes, 100)
    fmt.Printf("\n:: creating worker-pool\n")
    wp         := chase.WPCreate()
    fmt.Printf("\n:: worker pool has been created\n")

    fmt.Printf("\n:: start chasing ..\n")
    _ = chase.Listen(path1, messages, wp)
    fmt.Printf("\n:: chase has been started ..\n")
    evebridge.Handle(messages)
    //
}
//
//
//
//
//
