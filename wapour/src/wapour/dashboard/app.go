package dashboard
import . "wapour/dashboard/actions"
import . "wapour/dashboard/files"
import . "wapour/dashboard/hosts"

import "wapour/api/wengine"

import "github.com/gin-gonic/gin"
//import "html/template"

func ActionsView()(func (c *gin.Context)) {

    //template_name := "index.html"
    //data["content"] = template.HTML(Actions())
    return  func(c *gin.Context ){
        //c.HTML(200, template_name,  data )
        c.String(200,Actions())
    }
}

func ActionsJson()(func (c *gin.Context)) {

    rows := [][]string{ }
    api := wengine.GetApi("","","")
    _,action_list:=api.ActionsList()
    // 
    header := []string {"Name","Command"}
    for id := range action_list {
        action:=action_list[id]
        row  :=[]string{ action.Name, action.Command}
        rows = append( rows, row )

    }
    data_items:=[]gin.H{ gin.H{"data_type":"wapour-table","name":"Actions Table","id":"actions_table","title":"Actions Table","rows":rows, "header":header}}
    var data = gin.H{"status": "ok","data_items":data_items}



    //
    return  func(c *gin.Context ){
        //c.HTML(200, template_name,  data )
        //c.String(200,Actions())
        c.JSON(200, gin.H{"status": "ok","data":data})
    }
}


func FilesView()(func (c *gin.Context)) {
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

func HostsView()(func (c *gin.Context)) {
    return  func(c *gin.Context ){
        c.String(200, Hosts())
    }
}

func Overview()(func (c *gin.Context)) {

    template_name:="table.html"
    return  func(c *gin.Context ){
        c.HTML(200, template_name ,  gin.H{})
    }

}
