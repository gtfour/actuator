package rest
import "github.com/gin-gonic/gin"

var TOKEN_COOKIE_FIELD_NAME string = "USER_TOKEN"
var USERID_COOKIE_FIELD_NAME string = "USER_ID"


func GetTokenFromCookies(c *gin.Context)(token string,user string) {
    cookies:=c.Request.Cookies()
    for c := range cookies {
        cookie:=cookies[c]
        if cookie.Name == TOKEN_COOKIE_FIELD_NAME {
            token = cookie.Value
        }
        if cookie.Name == USERID_COOKIE_FIELD_NAME  {
            user = cookie.Value
        }

    }
    return user, token
}



func CheckPermissions(c *gin.Context)(bool) {
    return true
}
