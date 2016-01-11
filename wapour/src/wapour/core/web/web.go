package web
import "html/template"
import "io"

var templates_path = "/actuator/wapour/src/wapour/core/web/templates/*"
var templates = template.Must(template.ParseGlob(templates_path))


func RenderTemplate (template_name string, data interface{} )(rendered string) {

    var writer io.Writer
    _ = templates.ExecuteTemplate(writer, template_name, data)
    io.WriteString(writer, rendered)
    return rendered

}
