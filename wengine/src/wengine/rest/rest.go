package rest

import "github.com/gin-gonic/gin"
import "encoding/json"
import "fmt"
import "wengine/dusk"

func DashboardUsersList( data  gin.H, params ...[]string )(func (c *gin.Context)) {
    return  func(c *gin.Context ){
        c.String(200, "Hello")
    }
}

func GetUserById( data  gin.H, database dusk.Database ,c *gin.Context)(func(c *gin.Context)) {
    handler := func(c *gin.Context)( func(c *gin.Context) ) {

        id             := c.Query("userid")
        user,err_db    := database.GetUserById(id)
        b, err_marshal := json.Marshal(user)
        if err_db == nil && err_marshal == nil  {
            return func(c *gin.Context ){
                fmt.Printf("\nOk\n")
                c.String(200, string(b))
            }
        } else {
            return func(c *gin.Context ){
                fmt.Printf("\nError\n")
                status :="DB Error:"
                status += fmt.Sprintf("%v",err_db)
                status += "\nJSON Marshal Error:"
                status += fmt.Sprintf("%v",err_marshal)
                c.String(200, status)
            }
        }
    }
    return handler(c)
}




