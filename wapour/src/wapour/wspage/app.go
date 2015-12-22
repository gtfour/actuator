package wspage

import "github.com/gin-gonic/gin"

func WsPage( data  gin.H, params ...[]string )(func (c *gin.Context)) {

    template_name  := "ws.html"

    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }

}
