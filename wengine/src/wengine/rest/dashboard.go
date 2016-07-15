package rest
import "fmt"
//import "wengine/dusk"
import "github.com/gin-gonic/gin"

func GetDashboardData(data  gin.H)(func (c *gin.Context)) {
    return  func(c *gin.Context ){
        //if user_id,token_id,err:=auth.GetTokenFromCookies(c); auth.IsAuthorized(c) && err==nil {
            dashboardId        := c.Param("dashboardId")
            dashboardGroupId   := c.Param("dashboardGroupId")
            fmt.Printf("\nDashboard %v DashboardGroup %v\n",dashboardId,dashboardGroupId)
            //dashboard_data,_ := webclient.GetDashboardData(user_id, token_id, dashboardId)
            c.JSON(200, gin.H{"status": "ok","data":gin.H{}})
        //} else {
            //c.JSON(200, gin.H{"status": "error"})
        //}
    }
}

func SetDashboardData(data  gin.H)(func (c *gin.Context)) {
    return  func(c *gin.Context ){
        //if user_id,token_id,err:=auth.GetTokenFromCookies(c); auth.IsAuthorized(c) && err==nil {
            dashboardId        := c.PostForm("dashboardid")
            dashboardGroupId   := c.PostForm("dashboardgroupid")
            fmt.Printf("\nDashboard %v DashboardGroup %v\n",dashboardId,dashboardGroupId)
            //dashboard_data,_ := webclient.GetDashboardData(user_id, token_id, dashboardId)
            c.JSON(200, gin.H{"status": "ok","data":gin.H{}})
        //} else {
            //c.JSON(200, gin.H{"status": "error"})
        //}
    }
}

