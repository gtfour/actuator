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

    t = Table{TemplateName:"table"}
    return t

}

func (t *Table) AddAction ( ) {


}

func (table *Table) Render ()( fullfilled string ) {

    return RenderTemplate(table.TemplateName , table)

}
