package main

import "client/chase"
import "client/evebridge"

func main(){

    chaser := chase.NewChaser(0)
    app,_  := evebridge.MakeApp()
    app.AddMessageSource(chaser)
    chaser.Follow("/tmp/test")

}
