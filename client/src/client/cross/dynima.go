package cross

import "fmt"

type Dynima struct {
    //parsers
    Id              string
    ids             []int // id column number
    SourcePath      string
    SourceType      string
    header          []string
    data            [][]string
    filters         []string //FilterList
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


func CreateDynima(id string)(error) {

    fmt.Printf("Storage error %v\nNew dynima id %s\n",STORAGE_INSTANCE.Error,id)
    return nil

}

func EditDynima(id string)(error) {

    return nil


}

