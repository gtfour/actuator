package index

import "github.com/gin-gonic/gin"

func Index( context  gin.H, params ...[]string )(func (c *gin.Context)) {

    template_name:="index.html"
    
    return  func(c *gin.Context ){
        c.HTML(200, template_name ,  context )
    }

}


