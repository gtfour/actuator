package auth
import "errors"
import "github.com/gin-gonic/gin"
import "wapour/settings"
import "wapour/api/webclient"


func IsAuthorized( c *gin.Context , wrappers *[]*webclient.WengineWrapper)(bool) {
    token_id,user_id,err:=GetTokenFromCookies(c)
    if err!= nil {return false}
    w := webclient.FindWrapper(user_id,token_id,wrappers)
    if w == nil {return false} else { return true}
}

func GetTokenFromCookies(c *gin.Context)(token string,user string,err error) {
    cookies:=c.Request.Cookies()
    for c := range cookies {
        cookie:=cookies[c]
        if cookie.Name == settings.TOKEN_COOKIE_FIELD_NAME {
            token = cookie.Value
        }
        if cookie.Name == settings.USERID_COOKIE_FIELD_NAME  {
            user = cookie.Value
        }

    }
    if (user == "" || token == "") {  return token,user,errors.New("token or user was not found in cookie")  }
    return token,user,nil
}

