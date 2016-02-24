package rest
import "github.com/gin-gonic/gin"
import "wengine/dusk"


func DuskUserRoute( data  gin.H, database dusk.Database ) ( func (c *gin.Context) ) {


        return func (c *gin.Context)  {
            param:=c.Param("duskModuleName")
            token_id,user_id,_:=GetTokenFromCookies(c)
            authorized := database.TokenExists(user_id,token_id)
            switch {
                case authorized == false:
                    Unauthorized(c)
                case param == "get-user-by-id":
                    handler:=GetUserById( data, database, c )
                    handler(c)
                case param == "get-my-dashboards":
                    handler:=GetMyDashboards( data, database, c )
                    handler(c)
                case param == "get-all-users":
                    handler:=GetAllUsers( data, database, c )
                    handler(c)
            }
            }
}

func Unauthorized(c *gin.Context ) {
    c.String(401, "Unauthorized")
}
