package utah
import "github.com/gin-gonic/gin"
import "fmt"

var UTAH_PORT string = "35357"

func AuthRoute( data  gin.H ) ( func (c *gin.Context) ) {


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
