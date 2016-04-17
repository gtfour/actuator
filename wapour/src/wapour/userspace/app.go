package userspace
import "github.com/gin-gonic/gin"
import "wapour/api/webclient"
import "wapour/auth"
import "wapour/core/common"


func Index( data  gin.H, wrappers *[]*webclient.WengineWrapper,  params ...[]string )(func (c *gin.Context)) {

    template_name  := "index.html"
    return  func(c *gin.Context ){
        //if (auth.IsAuthorized(c,wrappers) && (token_id,user_id,err:=auth.GetTokenFromCookies(c); err==nil) )  {
        // thanks for postman from golang@cjr
        if token_id,user_id,err:=auth.GetTokenFromCookies(c); auth.IsAuthorized(c,wrappers) && err==nil {
            session_id,_:=common.GenId()
            dashboards:=webclient.GetUserDashboards(token_id,user_id,wrappers)
            navigaton_menu := CreateNavigationMenu(dashboards)
            data["navigation_items"] = navigaton_menu
            data["session_id"]       = session_id
            c.HTML(200, template_name,  data )
        } else {
            c.Redirect(302,"/auth/login")
        }
    }
}
