package cuda

type Dynima struct {
    //parsers    
    ids             []int // id column number
    SourcePath      string
    SourceType      string
    header          []string
    data            [][]string
    parsers         ParserList
    insert_template 
}


func (d *Dynima) BindParser (parser_name string)(error) {
    return nil
}

func (d *Dynima) UnbindParser (parser_name string)(error) {
    return nil
}

func (d *Dynima) RunParser (parser_name string)(error) {
    return nil
}

func (d *Dynima) Save ()(error) {
    return nil
}

func (d *Dynima) GetData ()(error) {
    return nil
}
