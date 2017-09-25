package userspace

import "fmt"
import "github.com/gin-gonic/gin"
import "wapour/api/webclient"
import "wapour/auth"
import "wapour/settings"
import "wapour/core/common"


func Index()(func (c *gin.Context)) {
    template_name  := "index.html"
    self_link      := "/userspace"
    return  func(c *gin.Context ){
        //if (auth.IsAuthorized(c,wrappers) && (user_id,token_id,err:=auth.GetTokenFromCookies(c); err==nil) )  {
        // thanks for postman from golang@cjr
        if user_id,token_id,err :=auth.GetTokenFromCookies(c); auth.IsAuthorized(c) && err==nil {
            session_id,_             := common.GenId()
            dashboards               := webclient.GetUserDashboards(user_id, token_id)
            navigaton_menu           := CreateNavigationMenu(dashboards)
            data                     := settings.APP_SETTINGS
            data["navigation_menu"]  =  navigaton_menu
            data["session_id"]       =  session_id
            data["websocket"]        =  "true"
            data["app_data_url"]     = settings.USERSPACE_DATA_URL
            c.HTML(200, template_name,  data )
        } else {
            c.Redirect(302,settings.SERVER_URL+"/auth/login?redirect_to="+self_link)
        }
    }
}

func UserspaceData()(func (c *gin.Context)) {
    //template_name  := "index.html"
    return  func(c *gin.Context ){
        if user_id,token_id,err:=auth.GetTokenFromCookies(c); auth.IsAuthorized(c) && err==nil {
            dashboards               := webclient.GetUserDashboards(user_id, token_id)
            data                     := settings.APP_SETTINGS
            data["navigation_menu"] =  CreateNavigationMenuJson(dashboards)
            //data["status"]           = "ok"
            c.JSON(200, gin.H{"status": "ok","data":data})
        } else {
            c.JSON(401, gin.H{"status": "Unauthorized","data":gin.H{}})
        }
    }
}


func GetDashboardData()(func (c *gin.Context)) {
    return  func(c *gin.Context ){
        if user_id,token_id,err:=auth.GetTokenFromCookies(c); auth.IsAuthorized(c) && err==nil {
            dashboardId        := c.Param("dashboardId")
            dashboardGroupId   := c.Param("dashboardGroupId")
            fmt.Printf("\nDashboard %v DashboardGroup %v\n",dashboardId,dashboardGroupId)
            dashboard_data,_ := webclient.GetDashboardData(user_id, token_id, dashboardId)
            c.JSON(200, gin.H{"status": "ok","data":dashboard_data})
        } else {
            c.JSON(200, gin.H{"status": "error"})
        }
    }
}
