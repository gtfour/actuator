package rest
import "errors"
import "github.com/gin-gonic/gin"

var TOKEN_COOKIE_FIELD_NAME string = "USER_TOKEN"
var USERID_COOKIE_FIELD_NAME string = "USER_ID"


func GetTokenFromCookies(c *gin.Context)(token string,user string,err error) {
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
    if (user == "" || token == "") {  return token,user,errors.New("token or user was not found in cookie")  }
    return token,user,nil
}



func CheckPermissions(c *gin.Context)(bool) {
    return true
}
