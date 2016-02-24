package main

import   "github.com/gin-gonic/gin"
import   "wengine/dusk"
import   "wengine/rest"
import . "wengine/settings"

func main() {

    app      := gin.Default()
    database := dusk.OpenDatabase( PrimaryDatabase, DBusername , DBpassword , DBhost, DBdbname )
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
