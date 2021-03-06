package index

import "net/http"
import "fmt"
import "github.com/gin-gonic/gin"
import "wapour/settings"
import "wapour/auth"
import "wapour/salvo"
import "wapour/api/webclient"
import "wapour/core/common"

var userstorage = salvo.UserStorageInstance


func Index()(func (c *gin.Context)) {

    template_name  := "index.html"
    navigaton_menu := GetNavigationMenu()
    data:=gin.H{"navigation_items":navigaton_menu,"static_url":settings.STATIC_URL, "app_data_url":settings.ADMIN_DATA_URL,"websocket":"false" }
    return  func(c *gin.Context ){
        if auth.IsAuthorized(c) == true { c.HTML(200, template_name,  data ) } else { c.Redirect(302,settings.SERVER_URL+"/auth/login"+"?redirect_to="+"/index") }
    }
}

func IndexData()(func (c *gin.Context)) {

    return  func(c *gin.Context ){
        if auth.IsAuthorized(c) == true  {/*
            menu_name:="Admin Dashboard"
            var navigation_menu = gin.H { "menu_name": menu_name , "items":items    }
            data_items:=[]gin.H{ gin.H{"data_type":"wapour-table","name":"Files Table","id":"files_table","title":"Files Table","rows":rows, "header":header}}
            var data = gin.H{"data_items":data_items}*/
            var data = gin.H{"navigation_menu":GetNavigationMenuJson()}
            c.JSON(200, gin.H{"status": "ok","data":data})
        } else {
            c.JSON(401, gin.H{"status": "Unauthorized","data":gin.H{}})
        }
}
}

func Redirect(url string)(func (c *gin.Context)) {
    return  func(c *gin.Context ){ c.Redirect(302,settings.SERVER_URL+url) }
}



func Login( ) (func (c *gin.Context)) {
    template_name := "login.html"
    // 
    // modify for running through nginx
    // old:
    // // server_addr   := settings.SERVER_ADDR
    // // server_proto  := settings.SERVER_PROTO
    // // server_port   := settings.SERVER_PORT
    // // post_url      := server_proto+"://"+server_addr+":"+server_port+"/auth/login"
    //
    post_url:="/auth/login"
    data          :=gin.H{"post_url":post_url, "static_url":settings.STATIC_URL }
    return  func(c *gin.Context ){
        fmt.Printf("Login:GetRequest:\n---\n%v\n---\n",c.Request)
        redirect_to := c.DefaultQuery("redirect_to", "/index")
        if common.IsIn(redirect_to, settings.ALLOWED_REDIRECTS)==false {
            redirect_to = "/index"
        }
        if auth.IsAuthorized(c) == true {
            redirectTo := settings.SERVER_URL+redirect_to
            fmt.Printf("\n=== client is already authorized\t\tredirecting to: %s\n", redirectTo)
            c.Redirect(302,redirectTo)
        } else {
            data["post_url"] = post_url+"?redirect_to="+redirect_to // redirects if login post will be success 
            c.HTML(200, template_name,  data )
        }
    }
}

func Logout() (func (c *gin.Context)) {
    return  func(c *gin.Context ){
        user_id,token_id,err:=auth.GetTokenFromCookies(c)
        wrapper:=userstorage.GetWrapper(user_id,token_id)
        if wrapper != nil {
            webclient.Disconnect(wrapper)
            userstorage.RemoveWrapper(user_id,token_id)
            fmt.Printf("Logout:: FindWrapper:: %v", userstorage.FindWrapper(user_id,token_id))
        }
        if err == nil {
           c.Redirect(302,settings.SERVER_URL+"/auth/login")
        } else {
            c.Redirect(302,settings.SERVER_URL+"/index")
        }
        //c.HTML(200, template_name,  data )
    }
}


func LoginPost () (func (c *gin.Context)) {
    return func(c *gin.Context) {
        username    := c.PostForm("username")
        password    := c.PostForm("password")
        redirect_to := c.DefaultQuery("redirect_to", "/index")
        if common.IsIn(redirect_to, settings.ALLOWED_REDIRECTS)==false {
            redirect_to = "/index"
        }
        w, err :=  webclient.Init(username, password)
        if err != nil { c.Redirect(302,settings.SERVER_URL+"/auth/login"+"?redirect_to="+redirect_to) } else {
            user := userstorage.FindWrapper(w.UserId, w.TokenId)
            if user == nil {
                userstorage.AddWrapper(w)
            }
            cookie_userid := &http.Cookie{Name:settings.USERID_COOKIE_FIELD_NAME, Value:w.UserId, Path:"/", Domain:settings.SERVER_ADDR }
            cookie_token  := &http.Cookie{Name:settings.TOKEN_COOKIE_FIELD_NAME,  Value:w.TokenId, Path:"/", Domain:settings.SERVER_ADDR }
            http.SetCookie(c.Writer, cookie_userid)
            http.SetCookie(c.Writer, cookie_token)


            redirectTo := settings.SERVER_URL+redirect_to
            fmt.Printf("\n=== login has been success \t\t redirecting to %s\n", redirectTo )

            c.Redirect(302,settings.SERVER_URL+redirect_to)
        }
    }
}
