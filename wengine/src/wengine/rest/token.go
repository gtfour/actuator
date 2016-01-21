package rest
import "github.com/gin-gonic/gin"

func CheckToken(c *gin.Context)(bool) {
    return true
}



func CheckPermissions(c *gin.Context)(bool) {
    return true
}
