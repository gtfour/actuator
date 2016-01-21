package rest

import "github.com/gin-gonic/gin"
import "encoding/json"
import "wengine/dusk"
import .  "wengine/core/common"

func DashboardUsersList( data  gin.H, params ...[]string )(func (c *gin.Context)) {
    return  func(c *gin.Context ){
        c.String(200, "Hello")
    }
}

func GetUserById( data  gin.H, database dusk.Database ,c *gin.Context)(func(c *gin.Context)) {
    handler := func(c *gin.Context)( func(c *gin.Context) ) {
        id             := c.PostForm("userid")
        user,err_db    := database.GetUserById(id)
        b, err_marshal := json.Marshal(user)
        GetTokenFromCookies(c)
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
    return handler(c)
}






