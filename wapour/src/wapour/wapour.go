package main
import "wapour/index"
import "wapour/ws"
import "wapour/wspage"
import "wapour/userspace"
import "wapour/dashboard"
//import "wapour/api/webclient"
import "wapour/settings"
import "github.com/gin-gonic/gin"


func main() {


    //users:=make([]*webclient.WengineWrapper,0)


    app:= gin.Default()
    app.LoadHTMLGlob("/actuator/wapour/src/wapour/templates/*")                   // load app  templates
    //app.LoadHTMLGlob("/actuator/actuator/wapour/src/wapour/core/web/templates/*") // load core templates

    app.Static("/static", settings.STATIC_DIR)
    app.GET("/index",     index.Index(     settings.APP_SETTINGS ))
    app.GET("/userspace", userspace.Index( settings.APP_SETTINGS ))
    app.GET("/wspage", wspage.WsPage(settings.APP_SETTINGS ))
    dashboard_app:=app.Group("/dashboard")
    {
        dashboard_app.GET("/actions",  dashboard.ActionsView( settings.APP_SETTINGS) )
        dashboard_app.GET("/files",    dashboard.FilesView(   settings.APP_SETTINGS) )
        dashboard_app.GET("/hosts",    dashboard.HostsView(   settings.APP_SETTINGS) )
        dashboard_app.GET("/overview", dashboard.Overview(    settings.APP_SETTINGS) )
    }
    auth_app:=app.Group("/auth")
    {
        auth_app.GET( "/login", index.Login(settings.APP_SETTINGS) )
        auth_app.POST("/login", index.LoginPost(settings.APP_SETTINGS) )
        auth_app.GET( "/logout" , index.Logout(settings.APP_SETTINGS)  )
    }
    server:=ws.NewServer("/entry")
    go server.Listen()
    app.GET("/entry", ws.WSserver(gin.H{}, server.WShandler))
    app.Run(":8090")
}
