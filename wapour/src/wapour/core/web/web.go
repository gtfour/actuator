package web
import "html/template"
import "io"
import "fmt"

var templates_path = "/actuator/wapour/src/wapour/core/web/templates/*"
var templates = template.Must(template.ParseGlob(templates_path))


func RenderTemplate (template_name string, data interface{} )(rendered string) {

    var writer io.Writer
    err:= templates.ExecuteTemplate(writer, template_name, data)
    fmt.Printf("\n-----Error: %v---\n",err)
    io.WriteString(writer, rendered)
    return rendered

}
