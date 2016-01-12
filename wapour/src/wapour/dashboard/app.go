package dashboard
import . "wapour/dashboard/actions"
import . "wapour/dashboard/files"
import "github.com/gin-gonic/gin"
import "html/template"

func ActionsView( data  gin.H, params ...[]string )(func (c *gin.Context)) {

    template_name := "index.html"
    data["content"] = template.HTML(Actions())


    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }

}

func FilesView( data  gin.H, params ...[]string )(func (c *gin.Context)) {

    template_name := "index.html"
    var content string
    content =  Files()
    content += Actions()
    data["content"] = template.HTML(content)


    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }

}



