package table
import . "wapour/core/web"

type Row  struct {

    Fields []string

}

type Table struct {

    Name         string
    Id           string
    Title        string
    HeaderFields Row
    FooterFields Row
    Rows         []Row
    Paginate     bool
    Search       bool
    Ajaxed       bool
    TemplateName string

}

func CreateTable()(t Table) {

    return t

}

func (t *Table) AddAction ( ) {


}

func (t *Table) Render (data interface {})  ( fullfilled string ) {

    return RenderTemplate(t.TemplateName , t)

}
