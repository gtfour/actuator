package cuda

import "errors"

var dup_name     = errors.New("error:filter with following name is already exist")
var name_is_none = errors.New("error:filter name wasn't specified")

var Filters = CreateFilterList()


type FilterList []Filter

func (fl FilterList) Append (new_filter Filter)(error) {

    if (new_filter.Name == "") { return name_is_none }
    for key := range fl {
        filter := fl[key]
        if ( new_filter.Name == filter.Name ) {
            return dup_name
        }
    }
    fl = append(fl, new_filter)
    return nil
}


type Filter struct {

    Name string
    Call func( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int)

}


func CreateFilterList ()(FilterList) {

    fl := make(FilterList, 0)
    //pl.Append(shitty_parser)
    return fl

}


