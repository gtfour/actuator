package rest
import "github.com/gin-gonic/gin"
import "wengine/dusk"


func DuskRoute( data  gin.H, database dusk.Database ) ( func (c *gin.Context) ) {

        return func (c *gin.Context)  {
            param:=c.Param("duskModuleName")
            switch {
                case param == "get-user-by-id":
                    handler:=GetUserById( data, database, c )
                    handler(c)
            }
            }
}
