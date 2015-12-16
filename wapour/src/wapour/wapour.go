package main
import "wapour/overview"
import "github.com/gin-gonic/gin"

var STATIC_DIR = "/actuator/wapour/static"
var STATIC_URL = "/static/main/"


func main() {

    app:= gin.Default()

    app.LoadHTMLGlob("/actuator/wapour/src/wapour/templates/*")

    app.Static("/static","/actuator/wapour/static")

    app.GET("/overview", overview.Overview( gin.H{"static_url":STATIC_URL}) )
    app.GET("/index", Index(                gin.H{"static_url":STATIC_URL}) )

    app.Run(":8090")

}

func Index( context  gin.H, params ...[]string )(func (c *gin.Context)) {

    template_name:="index.html"
    return  func(c *gin.Context ){
        c.HTML(200, template_name ,  context )
    }

}

