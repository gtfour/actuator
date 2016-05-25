package cuda

type Dynima struct {
    //parsers    
    Id              string
    ids             []int // id column number
    SourcePath      string
    SourceType      string
    header          []string
    data            [][]string
    filters         FilterList
    template        string
    data_indexes    [][]int
    delim_indexes   [][]int
}


func (d *Dynima) BindFilter (filter_name string)(error) {
    return nil
}

func (d *Dynima) UnbindFilter (filter_name string)(error) {
    return nil
}

//func (d *Dynima) RunFilter (filter_name string)(error) {
//    return nil
//}

func (d *Dynima) Save ()(error) {
    return nil
}

func (d *Dynima) GetData ()(error) {
    return nil
}

func (d *Dynima) SetTemplate()(error) {
    return nil
}

func (d *Dynima) SetSource (sourceType string, sourcePath string)(error) {
    return nil
}

