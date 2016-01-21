package rest
import "github.com/gin-gonic/gin"
import "wengine/dusk"


func DuskUserRoute( data  gin.H, database dusk.Database ) ( func (c *gin.Context) ) {

        return func (c *gin.Context)  {
            param:=c.Param("duskModuleName")
            switch {
                case param == "getuserbyid":
                    handler:=GetUserById( data, database, c )
                    handler(c)
            }
            }
}
