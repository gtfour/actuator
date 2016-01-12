package dashboard
import . "wapour/dashboard/actions"
import . "wapour/dashboard/files"
import . "wapour/dashboard/hosts"

import "github.com/gin-gonic/gin"
//import "html/template"

func ActionsView( data  gin.H, params ...[]string )(func (c *gin.Context)) {

    //template_name := "index.html"
    //data["content"] = template.HTML(Actions())
    return  func(c *gin.Context ){
        //c.HTML(200, template_name,  data )
        c.String(200,Actions())
    }
}

func FilesView( data  gin.H, params ...[]string )(func (c *gin.Context)) {
    //template_name := "index.html"
    //var content string
    //content =  Files()
    // content += Actions() : Example of creating several tables
    //data["content"] = template.HTML(content)
    return  func(c *gin.Context ){
        //c.HTML(200, template_name,  data )
        c.String(200, Files())
    }
}

func HostsView( data  gin.H, params ...[]string )(func (c *gin.Context)) {
    return  func(c *gin.Context ){
        c.String(200, Hosts())
    }
}

func Overview( data  gin.H, params ...[]string )(func (c *gin.Context)) {

    template_name:="table.html"
    return  func(c *gin.Context ){
        c.HTML(200, template_name ,  gin.H{})
    }

}
