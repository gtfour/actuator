package cuda

import "encoding/json"

var RESULT_TYPE_LINE    int = 7002
var RESULT_TYPE_SECTION int = 7004

type Result interface {
    ProceedTemplate([][]string)  string
    GetData        ()            ([]Line,error)
    GetType        ()            (typ int)
    GetJson        ()            ([]byte,error)
}


type Line struct {
    data_string_slice         []string    `json:"data_string_slice"`
    data_indexes              [][]int     `json:"data_indexes"`
    delim_indexes             [][]int     `json:"delim_indexes"`
    data                      [][]string  `json:"data"`
    template                  string      `json:"template"`
    template_data_size        int         `json:"template_data_size"`
}

type Section struct {
    lines                     []Line      `json:"lines"`
    template                  string      `json:"template"`
    template_data_size        int         `json:"template_data_size"`
    typer                     int         `json:"typer"`
}


func(l *Line)GetData()(lines []Line,err error){
    if (l!=nil){
        lines = make([]Line,   0)
        lines = append(lines, *l)
        return lines, nil
    } else {
        return nil, nilResultError
    }
}

func (l *Line)GetType()(int){
    return RESULT_TYPE_LINE
}

func (l *Line)GetJson()([]byte,error){
    if l == nil {return nil,nilResultError}
    line_byte,err := json.Marshal(l)
    if err != nil {
        return nil, err
    } else {
        return line_byte,nil
    }
}

func(s *Section)GetData()([]Line,error){
    if (s.lines!=nil){
        return s.lines, nil
    } else {
        return nil, nilResultError
    }
}

func (s *Section)GetType()(int){
    return RESULT_TYPE_SECTION
}

func (s *Section)GetJson()([]byte,error){
    if s == nil {return nil,nilResultError}
    section_byte,err := json.Marshal(s)
    if err != nil {
        return nil, err
    } else {
        return section_byte,nil
    }
}
