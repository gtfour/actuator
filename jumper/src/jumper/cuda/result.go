package cuda

var RESULT_TYPE_LINE    int = 7002
var RESULT_TYPE_SECTION int = 7004

type Result interface {
    ProceedTemplate([][]string)  string
    GetData        ()            ([]Line,error)
    GetType        ()            (typer int)
    GetAllJson     ()            (map[string]interface{},error)
}


type Line struct {
    data_string_slice         []string
    data_indexes              [][]int
    delim_indexes             [][]int
    data                      [][]string
    template                  string
    template_weird_data_size  int
}

type Section struct {
    lines                     []Line
    template                  string
    template_weird_data_size  int
    typer                     int
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
