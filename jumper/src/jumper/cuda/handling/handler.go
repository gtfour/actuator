package handling

import "strings"

import "jumper/cuda/result"
import "jumper/cuda/targets"
import "jumper/cuda/filtering"
import "jumper/cuda/analyze"
import "jumper/cuda/templating"

type Handler struct {
    //
    //  will be heavy structure with a lot of different fields
    //
    filters  filtering.FilterList
    target   *targets.Target
    single   bool
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
            h.single =  true
            r,e      := h.handleFile()
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
    baseSection       :=  result.NewSection( "" , result.SECTION_TYPE_BASE )
    // currentSectionId  :=  parentId
    //
    // baseSection contains data from whole file . will be appended to result.File.sections if any other sections won't be  found
    // sections are collecting while cycle below
    //
    var currentSection *result.Section
    // _,_ = baseSection, currentSection
    //
    currentSection           =   &baseSection
    //
    // currentSection will have baseSection address till any nested section won't be found inside that file 
    //
    defaultSectionBreaker    :=  func(string)(bool){ return false }
    writeToSectionInProgress :=  false
    sectionCouldBeNested     :=  false
    //
    //
    templateDataCounter      :=  0
    lineTemplate             :=  ""
    _                        = lineTemplate
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
        if defaultSectionBreaker(line) {
            parentSectionPointer:=currentSection.GetParentSectionPointer()
            if currentSection.GetType() != result.SECTION_TYPE_BASE && parentSectionPointer != nil {
                //
                // write current section to file
                var oldSection result.Section
                oldSection               = *currentSection
                //
                file.Append( oldSection )
                // switchback to parent section when current section will be close 
                currentSection = parentSectionPointer
                //
                //
            }
            if i!=len(lines)-1 {
                // means set writeToSectionInProgress to false just when this is not last line
                // if this last line we immediately  have to upload last section to sections slice
                writeToSectionInProgress = false
            }
            continue
        }
        //
        //
        //
        if section_type == analyze.NOT_SECTION {
            //
            //
            writeToSectionInProgress  =  true
            lineAsArray               := strings.Split(       line, ""   )
            delims,data               := analyze.GetIndexes( lineAsArray ) // as i remember GetIndexes just making base set of delims and data by  splitting line by spaces
            //
            // GetIndexes  make cause a bug or mistakes
            //
            // // ( lineAsArray, delims, data)(ndelims, ndata)
            //
            //
            for i := range h.filters {
                //
                filter := h.filters[i]
                if filter.Enabled {
                    new_delims, new_data  := filter.Call( lineAsArray, delims, data )
                    delims,     data      =  new_delims, new_data
                }
                //
            }
            //
            // gen template for the first line of the section
            //
            // templateDataCounter
            // if currentSection.Size() ==  0 {
            //    template, variableCounter    :=  templating.GenTemplate(lineAsArray, data)
            //    currentSection.LineTemplate  =   template
            // }
            tempTemplate, tempTemplateDataCounter := templating.GenTemplate( lineAsArray, data )
            if tempTemplateDataCounter > templateDataCounter {
                lineTemplate = tempTemplate
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
            // new section maybe found while writing to already opened section
            //
            current_section_type := currentSection.GetType()
            sectionCouldBeNested =  analyze.SectionCouldBeNested(current_section_type)
            //
            //
            if sectionCouldBeNested || currentSection.GetType() == result.SECTION_TYPE_BASE {
                //
                section_name           :=  line[section_name_indexes[0]:section_name_indexes[1]+1]
                childSection           :=  currentSection.NewChildSection( section_name , section_type )
                currentSection         =   &childSection
                newSectionBreaker      :=  GetSectionBreaker( line, section_name_indexes, section_tag_indexes, section_type )
                defaultSectionBreaker  =   newSectionBreaker
                //
            }
        }
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
    file.Append( baseSection ) // will append baseSection anyway
    if h.single {
        file.SetPath(target.GetPath())
    } else {
        file.SetPath(target.GetPathShort())
    }
    //
    return
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
