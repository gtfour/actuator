package main
import   "github.com/gin-gonic/gin"

var LOG_FILE string = "/tmp/connection.log"

func main(){

    GenData("hello")
    SplitIp("192.168.1.2")
    app      := gin.Default()
    app.POST("/check", Check())
    app.GET("/check", Check())
    app.POST("/insert",Insert())
    app.Run("0.0.0.0:9020")

}
