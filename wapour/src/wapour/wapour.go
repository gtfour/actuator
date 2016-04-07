package main
import "wapour/index"
import "wapour/ws"
import "wapour/wspage"
import "wapour/userspace"
import "wapour/dashboard"
import "wapour/api/webclient"
import "wapour/settings"
import "github.com/gin-gonic/gin"


func main() {


    users:=make([]*webclient.WengineWrapper,0)


    app:= gin.Default()
    app.LoadHTMLGlob("/actuator/wapour/src/wapour/templates/*")                   // load app  templates
    //app.LoadHTMLGlob("/actuator/actuator/wapour/src/wapour/core/web/templates/*") // load core templates

    app.Static("/static", settings.STATIC_DIR)
    app.GET("/index",     index.Index(     gin.H{ "static_url":settings.STATIC_URL, "ws_url":settings.WS_URL, "get_data_url":settings.GET_DATA_URL}, &users))
    app.GET("/userspace", userspace.Index( gin.H{ "static_url":settings.STATIC_URL, "ws_url":settings.WS_URL, "get_data_url":settings.GET_DATA_URL}, &users ))
    app.GET("/wspage", wspage.WsPage(gin.H{ "static_url":settings.STATIC_URL} ))
    dashboard_app:=app.Group("/dashboard")
    {
        dashboard_app.GET("/actions",  dashboard.ActionsView( gin.H{ "static_url":settings.STATIC_URL}) )
        dashboard_app.GET("/files",    dashboard.FilesView(   gin.H{ "static_url":settings.STATIC_URL}) )
        dashboard_app.GET("/hosts",    dashboard.HostsView(   gin.H{ "static_url":settings.STATIC_URL}) )
        dashboard_app.GET("/overview", dashboard.Overview(    gin.H{ "static_url":settings.STATIC_URL}) )
    }
    auth_app:=app.Group("/auth")
    {
        auth_app.GET( "/login", index.Login(gin.H{ "static_url":settings.STATIC_URL}) )
        auth_app.POST("/login", index.LoginPost(gin.H{ "static_url":settings.STATIC_URL}, &users) )
        auth_app.GET( "/logout" )
    }

    server:=ws.NewServer("/entry")
    go server.Listen()
    app.GET("/entry", ws.WSserver(gin.H{}, server.WShandler, &users ))
    app.Run(":8090")

}
