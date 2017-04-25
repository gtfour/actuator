package handling

import "strings"

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
        return targetIsNil
    }
    //
    //
    //
}


//

func NewHandler(config map[string]string)(h Handler){
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
    //
    target := h.target
    lines  := target.GetLines()
    //
    // passing empty name
    //
    baseSection :=  result.NewSection( "" , result.SECTION_TYPE_BASE )
    //
    //
    //
    // baseSection contains data from whole file . will be appended to result.File.sections if any other sections won't be  found
    // sections are collecting while cycle below
    //
    //
    //
    var currentSection *result.Section
    // _,_ = baseSection, currentSection
    //
    //
    //
    currentSection         =     &baseSection
    //
    // currentSection will have baseSection address till any nested section won't be found inside that file 
    //
    defaultSectionBreaker    :=  func(string)(bool){ return false }
    writeToSectionInProgress :=  false
    //
    //
    for i := range lines {
        //
        //
        //
        line := lines[i]
        //
        // check anyway if this string could be an section identifier 
        //
        section_name_indexes, section_tag_indexes, section_type := analyze.EscapeSection(line)
        //
        // _,_,_                                                 =   section_name_indexes, section_tag_indexes, section_type
        //
        if defaultSectionBreaker(line) { writeToSectionInProgress = false  ; continue }
        //
        //
        //
        if section_type == analyze.NOT_SECTION {
            //
            lineAsArray := strings.Split(       line, ""   )
            delims,data := analyze.GetIndexes( lineAsArray ) // as i remember GetIndexes just making base set of delims and data by  splitting line by spaces
            //
            // GetIndexes  make cause a bug or mistakes
            //
            // // ( lineAsArray, delims, data)(ndelims, ndata)
            //
            //
            for i:= range h.filters {
                //
                //
                //
                filter := h.filters[i]
                if filter.Enabled {
                    new_delims, new_data  := filter.Call( lineAsArray, delims, data )
                    delims,     data      =  new_delims, new_data
                }
                //
                //
                //
            }
            //
            //
            //
            selectedData := analyze.SelectDataByIndexes(lineAsArray, data)
            resultLine   := result.NewLine(selectedData, delims, data)
            //
            //
            //
            currentSection.Append(resultLine)
            //
            //
            //
        } else {
            //
            // New section may be found while reading file.
            // Let's add rule to prevent writing base section when this section is empty.
            //
            if currentSection.Size() > 0 {
                //
                //
                //
                writeToSectionInProgress = true
                var oldSection result.Section
                oldSection               = *currentSection
                file.Append(oldSection)
                //
                //
                //
            }
            //
            //
            //
            section_name           :=  line[section_name_indexes[0]:section_name_indexes[1]]
            childSection           :=  result.NewSection( section_name , section_type )
            currentSection         =   &childSection
            newSectionBreaker      :=  GetSectionBreaker( line, section_name_indexes, section_tag_indexes, section_type )
            defaultSectionBreaker  =   newSectionBreaker
            //
            //
            //
        }
        // 
        // 
        //
    }
    //
    // file.Size() is 0 when any nested section  has not been found
    //
    // check if section was not closed
    //
    if writeToSectionInProgress {
        var oldSection result.Section
        oldSection     = *currentSection
        file.Append( oldSection )
    }
    //
    if file.Size() == 0 { file.Append( baseSection ) }
    file.SetPath(target.GetPathShort())
    //
    //
    //
    return
    //
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
        //
        tgt        := nestedTargets[i]
        handler    := NewHandler(nil)
        handler.AddFilters( h.filters )
        handler.AddTargetPtr( &tgt )
        //
        resultFile,err :=  handler.handleFile()
        //
        if err == nil { directory.Append(resultFile) }
        //
    }
    //
    //
    return
    //
}
