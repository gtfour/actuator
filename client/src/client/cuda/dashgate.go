package cuda

var DASHGATE_TYPE_SINGLE      int = 10000
var DASHGATE_TYPE_SINGLE_FILE int = 10001
var DASHGATE_TYPE_GROUP       int = 10002

type Dashgate struct {
    Id               string
    SourcePath       string
    SourceType       string // file or dir or command
    //
    //
    FileBaseDir      string
    Dynimas          []*Dynima
    DefaultDynima    *Dynima
    FileExtension    string
}

func (d *Dashgate)PassDataToDynima(dynima_id string)(error){

    return nil

}


func (d *Dashgate)EditDefaultDynima(dynima_id string)(error){

    return nil

}

func (d *Dashgate)Update()(error) {

    return nil

}
