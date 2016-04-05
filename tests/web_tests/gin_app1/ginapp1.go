package main

import "github.com/gin-gonic/gin"

var STATIC_DIR           = "/actuator/wapour/static"
var STATIC_URL           = "/static/main/"
var LOCAL_STATIC_URL     = "/local-static/"


func main() {



    app:= gin.Default()
    app.LoadHTMLGlob("/actuator/tests/web_tests/gin_app1/templates/*")

    app.Static("/static","/actuator/wapour/static")
    app.Static("/local-static","/actuator/tests/web_tests/gin_app1/local_static")
    app.GET("/index",         Index( gin.H{"static_url":STATIC_URL,"local_static_url":LOCAL_STATIC_URL}))
    app.Run(":9010")

}

func Index( data  gin.H,  params ...[]string )(func (c *gin.Context)) {

    template_name  := "index.html"
    //navigaton_menu := GetNavigationMenu()
    //data["navigation_items"] = navigaton_menu
    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }
}
