package main
import "wapour/overview"
import "wapour/index"
//import "wapour/ws"
import "wapour/wspage"
import "github.com/gin-gonic/gin"

var STATIC_DIR = "/actuator/wapour/static"
var STATIC_URL = "/static/main/"


func main() {

    app:= gin.Default()

    app.LoadHTMLGlob("/actuator/wapour/src/wapour/templates/*")

    app.Static("/static","/actuator/wapour/static")

    app.GET("/overview", overview.Overview( gin.H{"static_url":STATIC_URL}) )
    //app.GET("/index",          index.Index( gin.H{"static_url":STATIC_URL, "navigation_items":[]string{"Events","Actions","Triggers"}}) )
    app.GET("/index",          index.Index( gin.H{"static_url":STATIC_URL}))
    //app.GET("/ws", ws.WS(gin.H{}))
    app.GET("/wspage", wspage.WsPage(gin.H{ "static_url":STATIC_URL} ))

    app.Run(":8090")

}
