package main

import "github.com/gin-gonic/gin"
import "wengine/core/utah"

func main() {

    app:= gin.Default()
    app.POST("/auth/:authModuleName",  utah.AuthRoute( gin.H{} ) )
    app.GET("/auth/:authModuleName",  utah.AuthRoute( gin.H{} ) )
    app.Run(":9000")



}
