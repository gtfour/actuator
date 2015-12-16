package overview
//package main

import "github.com/gin-gonic/gin"

func Overview( context  gin.H, params ...[]string )(func (c *gin.Context)) {

    template_name:="table.html"
    return  func(c *gin.Context ){
        c.HTML(200, template_name ,  gin.H{})
    }

}
