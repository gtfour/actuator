package web
import "html/template"
import "bytes"
import "fmt"

var templates_path = "/actuator/wapour/src/wapour/core/web/templates/*"
var templates = template.Must(template.ParseGlob(templates_path))


func RenderTemplate (template_name string, data interface{} )(rendered string) {

    var buf *bytes.Buffer = &bytes.Buffer{}
    //fmt.Printf("\n--- Data: %v ---\n--- TemplateName: %v ---\n",data,template_name)
    _= templates.ExecuteTemplate(buf, template_name, data)
    //fmt.Printf("\n-----Error: %v---\n",err)
    test:=buf.String()
    fmt.Printf("\n::: Rendered string:%s\n",test)
    return test

}
