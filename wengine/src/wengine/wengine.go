package main

import "github.com/gin-gonic/gin"
import "wengine/dusk"
import "wengine/rest"

func main() {

    app      := gin.Default()
    database := dusk.OpenDatabase("mongo","wengine","OpenStack123","127.0.0.1","wengine")
    defer database.Close()
    app.POST("/auth/:authModuleName",  rest.AuthRoute( gin.H{}, database ))
    app.GET("/auth/:authModuleName",   rest.AuthRoute( gin.H{}, database ))
    restapp:=app.Group("/rest")
    {
        restapp.POST("/user/:duskModuleName", rest.DuskUserRoute( gin.H{},  database ) )
        restapp.GET("/user/:duskModuleName",  rest.DuskUserRoute( gin.H{},  database ) )
    }
    app.Run(":9000")



}
