package rest
import "github.com/gin-gonic/gin"
import "wengine/dusk"

import "fmt"

var UTAH_PORT string = "35357"

func AuthRoute( data  gin.H, database dusk.Database) ( func (c *gin.Context) ) {
    return func (c *gin.Context) {
        param:=c.Param("authModuleName")
        fmt.Printf("\n--param: %s --\n",param)
        switch {
            case param == "login":
                fmt.Printf("\n--login--\n")
            case param == "logout":
                fmt.Printf("\n--logout--\n")
        }
    }
}

func LoginSuccess(c *gin.Context ) {
    c.String(200, "LoggedIn")
}
func LoginFailed(c *gin.Context ) {
    c.String(401, "Unauthorized")
}
