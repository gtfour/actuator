package index

import "github.com/gin-gonic/gin"

func Index( data  gin.H, params ...[]string )(func (c *gin.Context)) {

    template_name  := "index.html"
    navigaton_menu := GetNavigationMenu()


    data["navigation_items"] = navigaton_menu
    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }

}

func Login(data  gin.H, params ...[]string ) (func (c *gin.Context)) {

    template_name:="login.html"
    return  func(c *gin.Context ){
        c.HTML(200, template_name,  data )
    }

}
