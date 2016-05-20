package main

import   "github.com/gin-gonic/gin"
import   "wengine/dusk"
import   "wengine/rest"
import   "wengine/wsserver"
import   "wengine/settings"

func main() {

    app      := gin.Default()
    database := dusk.OpenDatabase( settings.PrimaryDatabase, settings.DBusername , settings.DBpassword , settings.DBhost, settings.DBdbname )
    defer database.Close()
    app.POST("/auth/:authModuleName",  rest.AuthRoute( gin.H{}, database ))
    app.GET("/auth/:authModuleName",   rest.AuthRoute( gin.H{}, database ))
    restapp:=app.Group("/rest")
    {
        restapp.POST("/user/:duskModuleName", rest.DuskUserRoute( gin.H{},  database ) )
        restapp.GET("/user/:duskModuleName",  rest.DuskUserRoute( gin.H{},  database ) )
        restapp.GET("/dashboard/get-dashboard-data/:dashboardGroupId/:dashboardId/",  rest.GetDashboardData( gin.H{},  database ) )
        restapp.POST("/dashboard/set-dashboard-data/", rest.SetDashboardData( gin.H{},  database ) )
    }
    app.GET(settings.WS_DATA_URL, wsserver.WebSocketHandle(gin.H{}))
    app.Run(":9000")



}
