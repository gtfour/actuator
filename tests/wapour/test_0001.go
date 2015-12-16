package main
import "github.com/gin-gonic/gin"

func main(){

    r:= gin.Default()
    r.GET("/hello/:name", func (c *gin.Context) {
        name:=c.Param("name")
        c.String(403,("hello "+name))})
    r.Run(":8090")
}
