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
    app.Static("/static", settings.STATIC_DIR)
    //app.LoadHTMLGlob("/actuator/actuator/wapour/src/wapour/core/web/templates/*") // load core templates
    //
    //
    app.GET("/",          index.Redirect("/userspace"))
    app.GET("/index",     index.Index())
    app.GET("/index-data",index.IndexData())
    app.GET("/userspace", userspace.Index())
    app.GET("/userspace-data",   userspace.UserspaceData())
    app.GET("/wspage",    wspage.WsPage(   settings.APP_SETTINGS ))
    //
    dashboard_app:=app.Group("/dashboard")
    {
        dashboard_app.GET("/actions",  dashboard.ActionsView())
        dashboard_app.GET("/files",    dashboard.FilesView())
        dashboard_app.GET("/hosts",    dashboard.HostsView())
        dashboard_app.GET("/overview", dashboard.Overview())

        dashboard_app.GET("/get-dashboard-data/:dashboardGroupId/:dashboardId/", userspace.GetDashboardData())
    }
    dashboard_json:=app.Group("/dashboard/json")
    {
        dashboard_json.GET("/actions",  dashboard.ActionsJson())
        dashboard_json.GET("/files",    dashboard.FilesJson())
        dashboard_json.GET("/hosts",    dashboard.HostsJson())

    }


    auth_app:=app.Group("/auth")
    {
        auth_app.GET( "/login", index.Login() )
        auth_app.POST("/login", index.LoginPost() )
        auth_app.GET( "/logout" , index.Logout()  )
    }
    server:=ws.NewServer("/entry")
    go server.Listen()
    app.GET("/entry", ws.WSserver(gin.H{}, server.WShandler))
    app.Run(":8090")
}
