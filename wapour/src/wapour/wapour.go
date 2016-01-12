package main
import "wapour/overview"
import "wapour/index"
import "wapour/ws"
import "wapour/wspage"
import "wapour/dashboard"
import "github.com/gin-gonic/gin"

var STATIC_DIR = "/actuator/wapour/static"
var STATIC_URL = "/static/main/"


func main() {

    app:= gin.Default()
    app.LoadHTMLGlob("/actuator/wapour/src/wapour/templates/*")                   // load app  templates
    //app.LoadHTMLGlob("/actuator/actuator/wapour/src/wapour/core/web/templates/*") // load core templates

    app.Static("/static","/actuator/wapour/static")
    app.GET("/overview", overview.Overview( gin.H{"static_url":STATIC_URL}) )
    //app.GET("/index",          index.Index( gin.H{"static_url":STATIC_URL, "navigation_items":[]string{"Events","Actions","Triggers"}}) )
    app.GET("/index",          index.Index( gin.H{"static_url":STATIC_URL}))
    app.GET("/wspage", wspage.WsPage(gin.H{ "static_url":STATIC_URL} ))
    app.GET("/dashboard/actions",dashboard.ActionsView(gin.H{ "static_url":STATIC_URL}) )

    server:=ws.NewServer("/entry")
    go server.Listen()
    app.GET("/entry", ws.WSserver(gin.H{}, server.WShandler ))

    app.Run(":8090")

}
