package rest
import "github.com/gin-gonic/gin"
import "wengine/dusk"
import "net/http"


type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func AuthRoute( data  gin.H, database dusk.Database) ( func (c *gin.Context) ) {
    return func (c *gin.Context) {
        param:=c.Param("authModuleName")
        switch {
            case param == "login":
                cred           := &Credentials{}
                err            := c.Bind(cred)
                if err != nil { c.JSON(401, gin.H{"status": "login_failed"}) } else {
                    username       := cred.Username
                    password       := cred.Password
                    userid,token,success := database.UserPasswordIsCorrect(username,password)
                    if ( success == false ) { c.JSON(401, gin.H{"status": "login_failed"}) } else {
                        cookie_userid := &http.Cookie{Name:USERID_COOKIE_FIELD_NAME, Value:userid}
                        cookie_token  := &http.Cookie{Name:TOKEN_COOKIE_FIELD_NAME,  Value:token}
                        http.SetCookie(c.Writer, cookie_userid)
                        http.SetCookie(c.Writer, cookie_token)
                        c.JSON(200, gin.H{"status": "login_success"})
                    }
                }
            case param == "logout":
                token_id,user_id,err := GetTokenFromCookies(c)
                if err!=nil {c.JSON(500, gin.H{"status": "not_logged_in"})}
                err                  = database.RemoveToken(token_id,user_id)
                if err!=nil { c.JSON(500, gin.H{"status": "logout_failed"})  } else {c.JSON(200, gin.H{"status": "logout_success"}) }
        }
    }
}
