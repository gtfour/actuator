package result

import "encoding/json"
import "jumper/common/gen"

type ResultsRaw []ResultRaw

type Result interface {
    //
    // // GetData        ()            ([]Line,error)
    //
    GetType        ()            (typ int)
    GetJson        ()            ([]byte,error)
    // // MarshalJSON()([]byte,error)
    //
    // // ProceedTemplate([][]string)  string
    //
}

type ResultSet struct {
    //
    Results    []Result
    // ResultsRaw []ResultRaw `json:"results"`
    //
}

type ResultRaw struct {
    Type       int             `json:"type"`
    Result     json.RawMessage `json:"result"`
}


type Line struct {
    //
    Data_string_slice         []string    `json:"data_string_slice"` // is it for lineAsArray ??? 
    delim_indexes             [][]int     `json:"delim_indexes"`
    data_indexes              [][]int     `json:"data_indexes"`
    Data                      [][]string  `json:"data"`
    // template                  string      `json:"template"`
    // template_data_size        int         `json:"template_data_size"`
    //
}



type Section struct {
    //
    //
    //
    Name                      string      `json:"name"`
    Id                        string      `json:"id"`
    ParentId                  string      `json:"parentId"`
    Typ                       int         `json:"typ"`
    Lines                     []Line      `json:"lines"`
    // LineTemplate           string      `json:"line_template"`
    // LineTemplate_data_size int         `json:"line_template_data_size"`
    //
    sectionTyp                int          // is it dup field ????
    parentSectionPointer      *Section
    //
    //
    //
}

type File struct {
    //
    Path                       string    `json:"path"`
    Sections                   []Section `json:"sections"`
    SectionTemplate            string    `json:"section_template"`
    SectionTemplate_data_size  int       `json:"section_template_data_size"`
    //
}

type Command struct {
    //
    Path                       string    `json:"path"`
    Sections                   []Section `json:"sections"`
    SectionTemplate            string    `json:"section_template"`
    SectionTemplate_data_size  int       `json:"section_template_data_size"`
    //
}


type Directory struct {
    //
    Path  string `json:"path"`
    Files []File `json:"files"` // have to change from lowercase to Uppercase because fields with lowecase don't visible after json.Marshal
    //
}

//
// Line methods
//

func(l *Line)GetData()(lines []Line,err error){
    //
    if ( l!=nil ){
        lines = make([]Line,   0)
        lines = append(lines, *l)
        return lines, nil
    } else {
        return nil, nilResultError
    }
    //
}

func (l *Line)GetType()(int){
    //
    return RESULT_TYPE_LINE
    //
}

func (l *Line)GetJson()([]byte,error){
    if l == nil { return nil,nilResultError }
    line_byte,err := json.Marshal(l)
    if err != nil {
        return nil, err
    } else {
        return line_byte,nil
    }
}

/*
func (l *Line)Append(line []string)(){
    if line != nil {
        l.Data = append(l.Data, line)
    }
}
*/
//
// Section methods
//

func(s *Section)GetData()([]Line,error){
    if (s.Lines!=nil){
        return s.Lines, nil
    } else {
        return nil, nilResultError
    }
}

func (s *Section)GetType()(int){
    //
    return s.Typ
    //
}

func (s *Section)GetId()(string){
    //
    return s.Id
    //
}

func (s *Section)SetParentId(parentId string)(){
    //
    s.ParentId = parentId
    //
}



func (s *Section)GetJson()([]byte,error){
    if s == nil { return nil,nilResultError }
    section_byte,err := json.Marshal(s)
    if err != nil {
        return nil, err
    } else {
        return section_byte,nil
    }
}

func (s *Section)Append(line Line)(){
    s.Lines = append( s.Lines, line )
}

func (s *Section)Size()(int){
    return len(s.Lines)
}

func (s *Section)GetParentSectionPointer()(*Section){
    return s.parentSectionPointer

}

func (s *Section)NewChildSection(name string, typ int)(childSection Section){
    //
    childSection.Typ                   =  typ
    childSection.Name                  =  name
    childSection.Id,_                  =  gen.GenId()
    childSection.ParentId              =  s.Id
    childSection.parentSectionPointer  =  s
    //
    return
}



//
// File methods
//

func(f *File)GetData()([]Section,error){
    if ( f.Sections != nil ){
        return f.Sections, nil
    } else {
        return nil, nilResultError
    }
}

func (f *File)GetType()(int){
    return RESULT_TYPE_FILE
}

func (f *File)GetJson()([]byte,error){
    if f == nil { return nil,nilResultError }
    file_byte,err := json.Marshal(f)
    if err != nil {
        return nil, err
    } else {
        return file_byte,nil
    }
}

func (f *File)Append(s Section)(){
    //
    // additional section check needs
    //
    f.Sections = append(f.Sections, s)
}

func (f *File)Size()(int){
    return len(f.Sections)
}

func (f *File)GetPath()(string){
    return f.Path
}

func (f *File)SetPath(path string)(){
    f.Path = path
}


//
// Directory methods
//

func(d *Directory)GetData()( []File,error ){
    if ( d.Files!=nil ){
        return d.Files, nil
    } else {
        return nil, nilResultError
    }
}

func (d *Directory)GetType()(int){
    return RESULT_TYPE_DIR
}

func (d *Directory)GetJson()( []byte,error ){
    //
    if d == nil { return nil, nilResultError }
    directory_byte,err := json.Marshal(d)
    if err != nil {
        return nil, err
    } else {
        return directory_byte, nil
    }
    //
}

func (d *Directory)Append(file File)(){
    d.Files = append(d.Files, file)
}


//
// Result Set methods
//

func(rs *ResultSet)GetData()([]Result,error ){
    if ( rs.Results != nil ){
        return rs.Results, nil
    } else {
        return nil, nilResultError
    }
}

func(rs *ResultSet)GetType()(int){
    return RESULT_TYPE_SET
}

func(rs *ResultSet)GetJson()([]byte,error){
    //// if rs == nil { return nil, nilResultError }
    //// result_set_byte,err := json.Marshal(rs)
    //// if err != nil {
    ////     return nil, err
    //// } else {
    ////    return result_set_byte, nil
    //// }
    var resultsRaw ResultsRaw
    for i := range rs.Results {
        r       := rs.Results[i]
        rtype   := r.GetType()
        rjs,err := r.GetJson()
        if err == nil {
            resultRaw     := ResultRaw{Type:rtype, Result:rjs}
            resultsRaw =  append(resultsRaw, resultRaw)
        }

    }
    resultsetByte,err := json.Marshal(resultsRaw)
    if err != nil {
        return nil, err
    } else {
        return resultsetByte, nil
    }
}


func (rs *ResultSet)Append(result Result)(){
    //
    rs.Results = append(rs.Results, result)
    //
}




//
// Common methods
//

func NewLine(lineAsArray []string, delims [][]int, data [][]int)(line Line){
    //
    line.Data_string_slice = lineAsArray
    line.delim_indexes     = delims
    line.data_indexes      = data
    //
    return
}

func NewSection(name string, typ int)(s Section){
    //
    s.Typ    = typ
    s.Name   = name
    s.Id,_   = gen.GenId()
    //
    return
}

//
//
//
func BlankResult(rtype int)(Result){
    //
    switch rtype {
        case RESULT_TYPE_LINE:
            var line Line
            return &line
        case RESULT_TYPE_SECTION:
            var section Section
            return &section
        case RESULT_TYPE_FILE:
            var file File
            return &file
        case RESULT_TYPE_DIRECTORY:
            var directory Directory
            return &directory
        default:
            return nil
    }
    //
}
//
//
//
