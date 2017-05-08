package main

import   "github.com/gin-gonic/gin"
import   "wengine/dusk"
import   "wengine/rest"
import   "wengine/wsserver"
import   "wengine/settings"

func main() {
    //
    //
    app      := gin.Default()
    database := dusk.DATABASE_INSTANCE
    defer database.Close()
    app.POST( "/auth/:authModuleName",  rest.AuthRoute(gin.H{}))
    app.GET(  "/auth/:authModuleName",  rest.AuthRoute(gin.H{}))
    restapp:=app.Group("/rest")
    {
        restapp.POST("/user/:duskModuleName", rest.DuskUserRoute(gin.H{}))
        restapp.GET("/user/:duskModuleName",  rest.DuskUserRoute(gin.H{}))
        restapp.GET("/dashboard/get-dashboard-data/:dashboardGroupId/:dashboardId/",  rest.GetDashboardData(gin.H{}))
        restapp.POST("/dashboard/set-dashboard-data/", rest.SetDashboardData(gin.H{}))
    }
    app.GET(settings.WS_DATA_URL, wsserver.WebSocketHandle(gin.H{}))
    // s.GetHandler()
    app.Run(":9000")
    //
    //
}
