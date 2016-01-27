package rest
import "github.com/gin-gonic/gin"
import "wengine/dusk"
import "fmt"


func DuskUserRoute( data  gin.H, database dusk.Database ) ( func (c *gin.Context) ) {


        return func (c *gin.Context)  {
            param:=c.Param("duskModuleName")
            username,token:=GetTokenFromCookies(c)
            fmt.Printf("\nname:%s token:%s\n",username,token)
            authorized := database.TokenExists(username,token)
            switch {
                case authorized == false:
                    Unauthorized(c)
                case param == "getuserbyid":
                    handler:=GetUserById( data, database, c )
                    handler(c)
            }
            }
}

func Unauthorized(c *gin.Context ) {
    c.String(401, "Unauthorized")
}
