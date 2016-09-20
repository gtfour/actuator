package main

import "fmt"
import  "github.com/gin-gonic/gin"

import "run"
import "fila"

var LOG_FILE string = "/tmp/connection.log"

func main(){

    fmt.Printf("\nServer Props:\n%v\n",run.Props)

    fila.GenData("hello")
    fila.SplitIp("192.168.1.2")
    app      := gin.Default()
    app.POST("/check",  fila.Check() )
    app.GET("/check",   fila.Check() )
    app.POST("/insert", fila.Insert())
    app.Run(run.Props["server_addr"])

}
