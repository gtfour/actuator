package index

import "net/http"
import "github.com/gin-gonic/gin"
import "wapour/settings"
import "wapour/api/webclient"
import "wapour/auth"


func Index( data  gin.H, wrappers *[]*webclient.WengineWrapper,  params ...[]string )(func (c *gin.Context)) {

    template_name  := "index.html"
    navigaton_menu := GetNavigationMenu()
    data["navigation_items"] = navigaton_menu
    return  func(c *gin.Context ){
        if auth.IsAuthorized(c,wrappers) == true { c.HTML(200, template_name,  data ) } else { c.Redirect(302,"/auth/login") }
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

func LoginPost ( data  gin.H, wrappers *[]*webclient.WengineWrapper, params ...[]string ) (func (c *gin.Context)) {
    return func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")
        w, err :=  webclient.Init(username, password)
        if err != nil { c.Redirect(302,settings.SERVER_URL+"/auth/login") } else {
            user := webclient.FindWrapper(w.UserId, w.TokenId, wrappers)
            if user == nil {
                webclient.AppendWrapper(wrappers, &w)
            }
            cookie_userid := &http.Cookie{Name:settings.USERID_COOKIE_FIELD_NAME, Value:w.UserId, Path:"/", Domain:settings.SERVER_ADDR }
            cookie_token  := &http.Cookie{Name:settings.TOKEN_COOKIE_FIELD_NAME,  Value:w.TokenId, Path:"/", Domain:settings.SERVER_ADDR }
            http.SetCookie(c.Writer, cookie_userid)
            http.SetCookie(c.Writer, cookie_token)

            c.Redirect(302,settings.SERVER_URL+"/index")
        }
    }
}
