package dashboard
import . "wapour/dashboard/actions"

func Actions( data  gin.H, params ...[]string )(func (c *gin.Context)) {

    actions_table := Actions()


    data["actions_table"] = actions_table
    
    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }

}
