package rest
import "github.com/gin-gonic/gin"
import "wengine/dusk"


func DuskRoute( data  gin.H, database dusk.Database ) ( func (c *gin.Context) ) {

        param:=c.Param("authModuleName")
        switch {
            case param == "get-user-by-id":
                return GetUserById(data, database)
            //case param == "logout":
            //
        }
}
