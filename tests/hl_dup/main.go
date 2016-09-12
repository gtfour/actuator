package main
import   "github.com/gin-gonic/gin"

var LOG_FILE string = "/tmp/connection.log"

func main(){

    GenData("hello")
    app      := gin.Default()
    //app.POST("/check", Check())
    app.GET("/check", Check())
    app.Run("127.0.0.1:9020")

}
