package rest

import "github.com/gin-gonic/gin"
import "encoding/json"
import "wengine/dusk"
import .  "wengine/core/common"

type UserJson struct {
    UserId  string `json:"user_id"`
    TokenId string `json:"token_id"`
}

func DashboardUsersList( data  gin.H, params ...[]string )(func (c *gin.Context)) {
    return  func(c *gin.Context ){
        c.String(200, "Hello")
    }
}

func GetUserById( data  gin.H, database dusk.Database ,c *gin.Context)(func(c *gin.Context)) {
    handler := func(c *gin.Context)( func(c *gin.Context) ) {
        user_json      := &UserJson{}
        err            := c.Bind(user_json)
        if err != nil {  return func(c *gin.Context ) {  c.JSON(401, gin.H{"status": "login_failed"}) }  } else {
            user,err_db    := database.GetUserById(user_json.UserId)
            b, err_marshal := json.Marshal(user)
            if err_db == nil && err_marshal == nil  {
                return func(c *gin.Context ){
                    c.String(200, string(b))
                }
            } else {
                return func(c *gin.Context ){
                    errors:=CombineErrors(err_db, err_marshal)
                    e,_:= json.Marshal(errors)
                    c.String(200, string(e))
                }
            }
        }
    }
    return handler(c)
}

func GetMyDashboards(data  gin.H, database dusk.Database ,c *gin.Context) (func(c *gin.Context)) {
    handler := func(c *gin.Context)( func(c *gin.Context) ) {
        user_json           := &UserJson{}
        err                 := c.Bind(user_json)
        if err != nil { return func(c *gin.Context ){   c.JSON(401, gin.H{"status": "login_failed"}) } } else {
            user,err_db    := database.GetUserById(user_json.UserId)
            _, err_marshal := json.Marshal(user)
            if err_db == nil && err_marshal == nil  {
                dashboard_list,err_dash := database.GetMyDashboardList(user_json.UserId,user_json.TokenId)
                if err_dash == nil {
                    dashboard_list_json,_:= json.Marshal(dashboard_list)
                    return func(c *gin.Context ){
                        //c.String(200, string(b))
                        c.JSON(200, gin.H{"status": "success","dashboard_list":dashboard_list_json})
                    }
                } else { return func(c *gin.Context ){ c.JSON(401, gin.H{"status": "failed"}) } }
            } else {
                return func(c *gin.Context ){
                    errors:=CombineErrors(err_db, err_marshal)
                    e,_:= json.Marshal(errors)
                    c.String(200, string(e))
                }
            }
        }
    }
    return handler(c)
}
