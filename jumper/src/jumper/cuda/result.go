package cuda

import "encoding/json"

var RESULT_TYPE_LINE       int = 7002
var RESULT_TYPE_SECTION    int = 7004
var RESULT_TYPE_FILE       int = 7006
var RESULT_TYPE_DIRECTORY  int = 7008
var RESULT_TYPE_DIR        int = 7008

type Result interface {
    //
    GetData        ()            ([]Line,error)
    GetType        ()            (typ int)
    GetJson        ()            ([]byte,error)
    //
    ProceedTemplate([][]string)  string
    //
}


type Line struct {
    //
    //
    data_string_slice         []string    `json:"data_string_slice"`
    data_indexes              [][]int     `json:"data_indexes"`
    delim_indexes             [][]int     `json:"delim_indexes"`
    data                      [][]string  `json:"data"`
    template                  string      `json:"template"`
    template_data_size        int         `json:"template_data_size"`
    //
    //
}

type Section struct {
    //
    //
    lines                     []Line      `json:"lines"`
    template                  string      `json:"template"`
    template_data_size        int         `json:"template_data_size"`
    typ                       int         `json:"typ"`
    //
    sectionTyp                int
    //
    //
}

type File struct {
    //
    path      string     `json:"path"`
    sections  []Section  `json:"sections"`
    //
}

type Directory struct {
    //
    path      string    `json:"path"`
    files     []File    `json:"files"`
    //
}

//
// Line methods
//

func(l *Line)GetData()(lines []Line,err error){
    //
    if (l!=nil){
        lines = make([]Line,   0)
        lines = append(lines, *l)
        return lines, nil
    } else {
        return nil, nilResultError
    }
    //
}

func (l *Line)GetType()(int){
    return RESULT_TYPE_LINE
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
//
//
//
