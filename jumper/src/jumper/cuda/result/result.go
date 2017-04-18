package result

import "encoding/json"

type Result interface {
    //
    // // GetData        ()            ([]Line,error)
    //
    GetType        ()            (typ int)
    GetJson        ()            ([]byte,error)
    // // ProceedTemplate([][]string)  string
    //
}

type ResultSet struct {
    results []Result
}


type Line struct {
    //
    data_string_slice         []string    `json:"data_string_slice"`
    data_indexes              [][]int     `json:"data_indexes"`
    delim_indexes             [][]int     `json:"delim_indexes"`
    data                      [][]string  `json:"data"`
    template                  string      `json:"template"`
    template_data_size        int         `json:"template_data_size"`
    //
}

type Section struct {
    //
    //
    name                      string      `json:"name"`
    typ                       int         `json:"typ"`
    lines                     []Line      `json:"lines"`
    template                  string      `json:"template"`
    template_data_size        int         `json:"template_data_size"`
    //
    sectionTyp                int // is it dup field ????
    //
    //
}

type File struct {
    //
    Path      string     `json:"path"`
    sections  []Section  `json:"sections"`
    //
}

type Command struct {
    //
    Path      string     `json:"path"`
    sections  []Section  `json:"sections"`
    //
}


type Directory struct {
    //
    Path      string    `json:"path"`
    files     []File    `json:"files"`
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
    //
    if l == nil { return nil,nilResultError }
    line_byte,err := json.Marshal(l)
    if err != nil {
        return nil, err
    } else {
        return line_byte,nil
    }
    //
}


//
// Section methods
//

func(s *Section)GetData()([]Line,error){
    if (s.lines!=nil){
        return s.lines, nil
    } else {
        return nil, nilResultError
    }
}

func (s *Section)GetType()(int){
    //
    return RESULT_TYPE_SECTION
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

//
// File methods
//

func(f *File)GetData()([]Section,error){
    if ( f.sections != nil ){
        return f.sections, nil
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

func (f *File)Append(section Section)(){
    f.sections = append(f.sections, section)
}

//
// Directory methods
//

func(d *Directory)GetData()( []File,error ){
    if ( d.files!=nil ){
        return d.files, nil
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
    d.files = append(d.files, file)
}


//
// Result Set methods
//

func(rs *ResultSet)GetData()([]Result,error ){
    if ( rs.results != nil ){
        return rs.results, nil
    } else {
        return nil, nilResultError
    }
}

func(rs *ResultSet)GetType()(int){
    return RESULT_TYPE_SET
}

func(rs *ResultSet)GetJson()( []byte,error ){
    //
    if rs == nil { return nil, nilResultError }
    result_set_byte,err := json.Marshal(rs)
    if err != nil {
        return nil, err
    } else {
        return result_set_byte, nil
    }
    //
}

func NewSection(name string, typ int)(s Section){ s.typ = typ ; s.name = name ; return }

//
//
//

func BlankResult(rtype int)(Result){
    //
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
    //
}
//
//
//
