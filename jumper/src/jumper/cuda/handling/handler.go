package handling

import "jumper/cuda/result"
import "jumper/cuda/targets"
import "jumper/cuda/filtering"

type Handler struct {
    //
    //  will be heavy structure with a lot of different fields
    //
    filters filtering.FilterList
    target  *targets.Target
    //
    //

    //
    //
    //
}

func(h *Handler)AddFilters(filterList filtering.FilterList)(error){
    //
    //
    //
    if ( filterList != nil ) {
        h.filters = filterList
        return nil
    } else {
        return filterListIsNil
    }
    //
    //
    //
}

// AddTargetPtr


func(h *Handler)AddTargetPtr(target *targets.Target)(error){
    //
    //
    //
    if ( target != nil ) {
        h.target = target
        return nil
    } else {
        return filterListIsNil
    }
    //
    //
    //
}


//

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
