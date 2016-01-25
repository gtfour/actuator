package rest
import "github.com/gin-gonic/gin"
import "wengine/dusk"
import "net/http"


var UTAH_PORT string = "35357"

func AuthRoute( data  gin.H, database dusk.Database) ( func (c *gin.Context) ) {
    return func (c *gin.Context) {
        param:=c.Param("authModuleName")
        switch {
            case param == "login":
                username       := c.PostForm("username")
                password       := c.PostForm("password")
                userid,token,success := database.UserPasswordIsCorrect(username,password)
                if ( success == false ) { LoginFailed(c) } else {
                    cookie_userid := &http.Cookie{Name:USERID_COOKIE_FIELD_NAME, Value:userid}
                    cookie_token  := &http.Cookie{Name:TOKEN_COOKIE_FIELD_NAME,  Value:token}
                    http.SetCookie(c.Writer, cookie_userid)
                    http.SetCookie(c.Writer, cookie_token)
                    c.String(200, "LoggedIn")
                }
            case param == "logout":
                user_id, token_id := GetTokenFromCookies(c)
                err:=database.RemoveToken(token_id,user_id)
                if err!=nil { c.String(200, "LogoutFailed")  } else { c.String(200, "LoggedOut") }
        }
    }
}

func LoginSuccess(c *gin.Context ) {
    c.String(200, "LoggedIn")
}
func LogoutSuccess(c *gin.Context ) {
    c.String(200, "LoggedOut")
}
func LoginFailed(c *gin.Context ) {
    c.String(401, "Unauthorized")
}
