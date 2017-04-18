package handling

import "jumper/cuda/result"
import "jumper/cuda/targets"
import "jumper/cuda/filtering"
import "jumper/cuda/analyze"

type Handler struct {
    //
    //  will be heavy structure with a lot of different fields
    //
    filters  filtering.FilterList
    target   *targets.Target
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

func(h *Handler)Handle()(result.Result, error){
    //
    //
    //
    switch target_type:=h.target.GetType();target_type {
        //
        //
        case targets.TARGET_LINE:
            r,e := h.handleLine()
            return &r,e
        case targets.TARGET_FILE:
            r,e := h.handleFile()
            return &r,e
        case targets.TARGET_DIR:
            r,e := h.handleDirectory()
            return &r,e
        default:
            return nil, targetTypeUndefined
        //
        //
    }
    //
    //
    //
}




func(h *Handler)handleLine()(line result.Line, err error ){
    //
    return
    //
}

func(h *Handler)handleFile()(file result.File, err error ){
    //
    //
    target      :=  h.target
    lines       :=  target.GetLines()
    baseSection :=  result.NewSection( "", result.SECTION_TYPE_BASE )
    //
    //
    var currentSection result.Section
    _,_ = baseSection, currentSection
    //
    //
    for i := range lines {
        line := lines[i]
        //
        // check anyway if this string could be an section identifier 
        //
        section_name_indexes, section_tag_indexes, section_type := analyze.EscapeSection(line)
        _,_,_ = section_name_indexes, section_tag_indexes, section_type
        // 
        //
        //
    }
    return
    //
    //
}

func(h *Handler)handleDirectory()(directory result.Directory,err error ){
    //
    //
    target         :=  h.target
    nestedTargets  :=  target.GetNestedTargets()
    //
    //
    directory.Path = target.GetPath()
    for i := range nestedTargets {
        tgt        := nestedTargets[i]
        handler    := NewHandler(nil)
        handler.AddFilters(h.filters)
        handler.AddTargetPtr(&tgt)
        resultFile,err :=  handler.handleFile()
        if err == nil { directory.Append(resultFile) }
    }
    return
    //
}
