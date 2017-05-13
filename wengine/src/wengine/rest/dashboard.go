package rest
import "fmt"
//import "wengine/dusk"
import "wengine/wsserver"
import "github.com/gin-gonic/gin"

func AddDashboard(data  gin.H)(func (c *gin.Context)) {
    return  func( c *gin.Context ) {
        //
        // temporary handler just for fun :)
        //
        dashboardName   := c.PostForm("dashboardName")
        sourceType      := c.PostForm("sourceType")
        sourcePath      := c.PostForm("sourcePath")
        clientName      := c.PostForm("clientName")
        fmt.Printf("\n<|<| Add dashboard handler: dashboardName : %v sourceType : %v sourcePath : %v clientName : %s |>|>\n", dashboardName, sourceType, sourcePath, clientName)
        //
        //
        //
        wsServer           := wsserver.WebSocketServerWeb
        client             := wsServer.GetClientByName( clientName )
        clientHasBeenFound := false
        fmt.Printf("\n::-- checking websocket clients --::\n")
        if client != nil { clientHasBeenFound = true }
        fmt.Printf("\n::- All clients list -::\n")
        all_ws_clients := wsServer.GetClients()
        fmt.Printf("\n%v\n--- --- ---\n", all_ws_clients)
        // 
        // json-response
        //
        wsInfoResponse := gin.H{ "clientName":clientName, "isFound":clientHasBeenFound }
        c.JSON(200, gin.H{"status": "ok","data":wsInfoResponse})
        //
        //
        //
    }
}

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

func DashboardRoute( data  gin.H ) ( func (c *gin.Context) ) {
        return func (c *gin.Context)  {
            param:=c.Param("duskModuleName")
            token_id,user_id,_:=GetTokenFromCookies(c)
            authorized := database.TokenExists(user_id,token_id)
            switch {
                case authorized == false:
                    Unauthorized(c)
                case param == "get-user-by-id":
                    handler:=GetUserById( data, c )
                    handler(c)
                case param == "get-my-dashboards":
                    handler:=GetMyDashboards( data, c )
                    handler(c)
                case param == "get-all-users":
                    handler:=GetAllUsers( data, c )
                    handler(c)
            }
            }
}
