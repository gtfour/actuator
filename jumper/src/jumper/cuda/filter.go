package cuda

import "errors"
// import "jumper/cuda/filtering"
// import "client/cuda/custom"

var dup_name         =  errors.New("error:filter with following name is already exist")
var name_is_none     =  errors.New("error:filter name wasn't specified")

var GetCustonFilters =  func()(fl FilterList){ return fl  } ;
var DefaultFilters   =  CreateDefaultFilterList()


type FilterList []Filter

type Filter struct {
    Name     string
    Enabled  bool
    Call     func( lineAsArray []string , delims [][]int , data [][]int)(ndelims [][]int , ndata [][]int)
}

func (fl *FilterList)Append(new_filter Filter)(error) {
    if (new_filter.Name == "") { return name_is_none }
    for key := range (*fl) {
        filter := (*fl)[key]
        if ( new_filter.Name == filter.Name ) {
            return dup_name
        }
    }
    (*fl) = append((*fl), new_filter)
    return nil
}

func (fl *FilterList)RemoveByName(filter_name string)(error) {
    //
    if ( filter_name == "") { return name_is_none }
    //
    var selected_filters_indexes []int
    for key := range (*fl) {
        filter := (*fl)[key]
        if ( filter_name == filter.Name ) {
            selected_filters_indexes = append( selected_filters_indexes, key )
        }
    }
    //
    for i:= range selected_filters_indexes {
        indexToRemove:=selected_filters_indexes[i]
        (*fl) = append((*fl)[:indexToRemove], (*fl)[indexToRemove+1:]...)
    }
    //
    return nil
}



func CreateDefaultFilterList ()(fl FilterList) {
    //
    fl = make(FilterList,0)
    //
    var url_filter      = Filter{ Name:"url_filter",      Call:UrlFilter,      Enabled:true }
    var path_filter     = Filter{ Name:"path_filter",     Call:PathFilter,     Enabled:true }
    var brackets_filter = Filter{ Name:"brackets_filter", Call:BracketsFilter, Enabled:true }
    var quotes_filter   = Filter{ Name:"quotes_filter",   Call:QuotesFilter,   Enabled:true }
    //
    fl.Append(url_filter)
    fl.Append(path_filter)
    fl.Append(brackets_filter)
    fl.Append(quotes_filter)
    //
    var custom_filters = GetCustonFilters()
    for i:= range custom_filters {
        filter:=custom_filters[i]
        fl.Append(filter)
    }
    return fl
    //
}
