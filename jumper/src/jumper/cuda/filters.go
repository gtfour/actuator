package cuda

import "errors"
//import "client/cuda/custom"

var dup_name     = errors.New("error:filter with following name is already exist")
var name_is_none = errors.New("error:filter name wasn't specified")

var GetCustonFilters = func()(fl FilterList){ return fl  } ;
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

    Name     string
    Enabled  bool
    Call     func( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int)

}


func CreateFilterList ()(FilterList) {

    fl := make(FilterList,0)

    var url_filter      = Filter{Name:"url_filter",      Call:UrlFilter,      Enabled:true}
    var path_filter     = Filter{Name:"path_filter",     Call:PathFilter,     Enabled:true}
    var brackets_filter = Filter{Name:"brackets_filter", Call:BracketsFilter, Enabled:true}
    var quotes_filter   = Filter{Name:"quotes_filter",   Call:QuotesFilter,   Enabled:true}


    fl.Append(url_filter)
    fl.Append(path_filter)
    fl.Append(brackets_filter)
    fl.Append(quotes_filter)

    var custom_filters = GetCustonFilters()
    for i:= range custom_filters {
        filter:=custom_filters[i]
        fl.Append(filter)
    }

    return fl

}


