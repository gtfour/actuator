package table

type Row  struct {

    Fields []string

}

type Table struct {

    Name         string
    HeaderFields []string
    FooterFields []string
    Rows         []Row
    Paginate     bool
    Search       bool

}

func CreateTable()(t Table) {

    return t

}

func (t *Table) AddAction ( ) {


}

