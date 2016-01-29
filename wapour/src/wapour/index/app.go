package index

import "net/http"
import "errors"
import "fmt"
import "github.com/gin-gonic/gin"
import "wapour/settings"
import "wapour/api/webclient"

var TOKEN_COOKIE_FIELD_NAME string = "USER_TOKEN"
var USERID_COOKIE_FIELD_NAME string = "USER_ID"


func Index( data  gin.H, wrappers []*webclient.WengineWrapper,  params ...[]string )(func (c *gin.Context)) {

    template_name  := "index.html"
    navigaton_menu := GetNavigationMenu()
    data["navigation_items"] = navigaton_menu
    return  func(c *gin.Context ){
        if IsAuthorized(c,wrappers) == true { c.HTML(200, template_name,  data ) } else { c.Redirect(302,"/auth/login") }
    }
}

func Login(data  gin.H, params ...[]string ) (func (c *gin.Context)) {
    template_name    := "login.html"
    server_addr      := settings.SERVER_ADDR
    server_proto     := settings.SERVER_PROTO
    server_port      := settings.SERVER_PORT
    post_url         := server_proto+"://"+server_addr+":"+server_port+"/auth/login"
    data["post_url"] = post_url
    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }
}

func LoginPost ( data  gin.H, wrappers []*webclient.WengineWrapper, params ...[]string ) (func (c *gin.Context)) {
    return func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")
        w, err :=  webclient.Init(username, password)
        if err != nil { c.Redirect(302,"/auth/login") } else {
            if user := webclient.FindWrapper(w.UserId, w.TokenId, wrappers); user == nil {
                webclient.AppendWrapper(wrappers, &w)
            }
            cookie_userid := &http.Cookie{Name:USERID_COOKIE_FIELD_NAME, Value:w.UserId}
            cookie_token  := &http.Cookie{Name:TOKEN_COOKIE_FIELD_NAME,  Value:w.TokenId}
            http.SetCookie(c.Writer, cookie_userid)
            http.SetCookie(c.Writer, cookie_token)
            c.Redirect(302,"/index")
        }
    }
}


func IsAuthorized( c *gin.Context , wrappers []*webclient.WengineWrapper)(bool) {
    token_id,user_id,err:=GetTokenFromCookies(c)
    if err!= nil {return false}
    w := webclient.FindWrapper(user_id,token_id,wrappers)
    if w == nil {return false} else { return true}
}

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
