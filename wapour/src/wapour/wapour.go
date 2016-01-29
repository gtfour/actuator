package main
import "wapour/index"
import "wapour/ws"
import "wapour/wspage"
import "wapour/dashboard"
import "wapour/api/webclient"
import "github.com/gin-gonic/gin"

var STATIC_DIR           = "/actuator/wapour/static"
var STATIC_URL           = "/static/main/"


func main() {


    users:=make([]*webclient.WengineWrapper,0)

    app:= gin.Default()
    app.LoadHTMLGlob("/actuator/wapour/src/wapour/templates/*")                   // load app  templates
    //app.LoadHTMLGlob("/actuator/actuator/wapour/src/wapour/core/web/templates/*") // load core templates

    app.Static("/static","/actuator/wapour/static")
    app.GET("/index",          index.Index( gin.H{"static_url":STATIC_URL}, users))
    app.GET("/wspage", wspage.WsPage(gin.H{ "static_url":STATIC_URL} ))
    dashboard_app:=app.Group("/dashboard")
    {
        dashboard_app.GET("/actions",  dashboard.ActionsView( gin.H{ "static_url":STATIC_URL}) )
        dashboard_app.GET("/files",    dashboard.FilesView(   gin.H{ "static_url":STATIC_URL}) )
        dashboard_app.GET("/hosts",    dashboard.HostsView(   gin.H{ "static_url":STATIC_URL}) )
        dashboard_app.GET("/overview", dashboard.Overview(    gin.H{ "static_url":STATIC_URL}) )
    }
    auth_app:=app.Group("/auth")
    {
        auth_app.GET( "/login", index.Login(gin.H{ "static_url":STATIC_URL}) )
        auth_app.POST("/login", index.LoginPost(gin.H{ "static_url":STATIC_URL}, users) )
        auth_app.GET( "/logout" )
    }

    server:=ws.NewServer("/entry")
    go server.Listen()
    app.GET("/entry", ws.WSserver(gin.H{}, server.WShandler ))

    app.Run(":8090")

}
