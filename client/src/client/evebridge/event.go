package evebridge

import "time"
import "fmt"

var LOG_CHANNEL_TIMEOUT_MS  time.Duration  = 1000

const (
      INITIALIZED  =  0 // initialized
      CREATED      =  1
      MODIFIED     =  2
      REMOVED      =  3)

type Event struct {

    Date string
    Path string
    Type string

}

type CompNotes struct {

    Path  string
    State int8
    List  []CompNote


}

type CompNote struct {

    Field    string
    Before   string
    After    string


}

func Handle(messages chan CompNotes )(err error) {

    for {

        select{
            case message:=<-messages:
                fmt.Println(message)

            default:
                time.Sleep( LOG_CHANNEL_TIMEOUT_MS  * time.Millisecond )
                fmt.Println("No messages")

        }

    }

}

