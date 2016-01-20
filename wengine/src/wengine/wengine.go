package main

import "github.com/gin-gonic/gin"
import "wengine/core/utah"
import "wengine/dusk"

func main() {

    app      := gin.Default()
    database := dusk.OpenDatabase("mongo","wengine","OpenStack123","127.0.0.1","wengine")
    app.POST("/auth/:authModuleName",  utah.AuthRoute( gin.H{} ) )
    app.GET("/auth/:authModuleName",  utah.AuthRoute( gin.H{} ) )
    app.POST("/dusk/:duskModuleName", utah.DuskRoute( gin.H{},  database ) )
    app.GET("/dusk/:duskModuleName",  utah.DuskRoute( gin.H{},  database ) )
    app.Run(":9000")



}
