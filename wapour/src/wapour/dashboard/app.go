package dashboard
import . "wapour/dashboard/actions"
import "github.com/gin-gonic/gin"
import "html/template"

func ActionsView( data  gin.H, params ...[]string )(func (c *gin.Context)) {

    actions_table := Actions()
    template_name := "index.html"
    data["content"] = template.HTML(actions_table)


    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }

}
