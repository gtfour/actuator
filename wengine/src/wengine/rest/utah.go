package rest
import "github.com/gin-gonic/gin"
import "wengine/dusk"
import "net/http"

import "fmt"

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
                    LoginSuccess(c)
                }
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
