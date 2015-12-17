package index

import "github.com/gin-gonic/gin"

type SubItem struct {

    Name string
    Url  string


}

type Item struct {
    Name     string
    SubItems []SubItem
}

func Index( data  gin.H, params ...[]string )(func (c *gin.Context)) {

    template_name:="index.html"
    sub_logs_list := SubItem{Name:"list",Url:"#logs/list"}
    sub_logs_add  := SubItem{Name:"add",Url:"#logs/add"}
    navigaton_menu:= []Item { Item { Name:"Logs" , SubItems:[]SubItem {sub_logs_list, sub_logs_add}  }}


    data["navigation_items"] = navigaton_menu
    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }

}
