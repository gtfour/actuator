package index

import "github.com/gin-gonic/gin"
import "wapour/settings"
import "wapour/api/webclient"
import "net/http"

var TOKEN_COOKIE_FIELD_NAME string = "USER_TOKEN"
var USERID_COOKIE_FIELD_NAME string = "USER_ID"


func Index( data  gin.H, params ...[]string )(func (c *gin.Context)) {

    template_name  := "index.html"
    navigaton_menu := GetNavigationMenu()


    data["navigation_items"] = navigaton_menu
    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }

}

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func Login(data  gin.H, params ...[]string ) (func (c *gin.Context)) {
    template_name    := "login.html"
    server_addr      := settings.SERVER_ADDR
    server_proto     := settings.SERVER_PROTO
    post_url         := server_proto+"://"+server_addr+"/auth/login"
    data["post_url"] = post_url
    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }
}

func LoginPost ( data  gin.H, wrappers []*webclient.WengineWrapper, params ...[]string ) (func (c *gin.Context)) {

    return func(c *gin.Context) {
        cred           := &Credentials{}
        err            := c.Bind(cred)
        if err != nil { Login(gin.H{"message":"Login/Password is invalid"}) }
        w, err :=  webclient.Init(cred.Username, cred.Password)
        if err != nil { Login(gin.H{"message":"Login/Password is invalid"}) }
        if user := webclient.FindWrapper(w.UserId, w.TokenId, wrappers); user == nil {
            webclient.AppendWrapper(wrappers, &w)
        }
        cookie_userid := &http.Cookie{Name:USERID_COOKIE_FIELD_NAME, Value:w.UserId}
        cookie_token  := &http.Cookie{Name:TOKEN_COOKIE_FIELD_NAME,  Value:w.TokenId}
        http.SetCookie(c.Writer, cookie_userid)
        http.SetCookie(c.Writer, cookie_token)
        Index(gin.H{"username":cred.Username})
    }
}
