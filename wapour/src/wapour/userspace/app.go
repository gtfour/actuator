package userspace
import "github.com/gin-gonic/gin"
import "wapour/api/webclient"
import "wapour/auth"
import "wapour/core/common"


func Index( data  gin.H, params ...[]string )(func (c *gin.Context)) {

    template_name  := "index.html"
    return  func(c *gin.Context ){
        //if (auth.IsAuthorized(c,wrappers) && (user_id,token_id,err:=auth.GetTokenFromCookies(c); err==nil) )  {
        // thanks for postman from golang@cjr
        if user_id,token_id,err:=auth.GetTokenFromCookies(c); auth.IsAuthorized(c) && err==nil {
            session_id,_:=common.GenId()
            dashboards  :=webclient.GetUserDashboards(token_id,user_id)
            navigaton_menu := CreateNavigationMenu(dashboards)
            data["navigation_items"] = navigaton_menu
            data["session_id"]       = session_id
            c.HTML(200, template_name,  data )
        } else {
            c.Redirect(302,"/auth/login")
        }
    }
}
