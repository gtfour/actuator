package main
import   "github.com/gin-gonic/gin"


func Insert()(func (c *gin.Context)) {
    return func (c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    }
}



