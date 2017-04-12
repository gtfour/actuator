package handling

import "jumper/cuda/result"
import "jumper/cuda/targets"
import "jumper/cuda/filtering"

type Handler struct {
    //
    filterList filtering.FilterList
    //
}

func(h *Handler)AddFilters(filterList filtering.FilterList)(error) {
    //
    //
    //
    if ( filterList != nil ) {
        h.filterList = filterList
        return nil
    } else {
        return filterListIsNil
    }
    //
    //
    //
}


func NewHandler(config map[string]string)(h *Handler){
    // 
    return h
    //
}

func(h *Handler)Handle(t targets.Target)(r result.Result){
    //
    return r
    //
}
