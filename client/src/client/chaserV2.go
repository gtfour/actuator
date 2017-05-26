package main

import "fmt"
import "client/chase"
import "client/evebridge"

func main(){

    chaser      := chase.NewChaser(0)
    chaser.Follow("/tmp/test")
    app,_       := evebridge.MakeApp()
    app.AddMessageSource(chaser)

}
