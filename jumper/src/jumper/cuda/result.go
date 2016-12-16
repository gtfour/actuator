package cuda

var RESULT_TYPE_LINE    int = 7002
var RESULT_TYPE_SECTION int = 7004

type Result interface {
    ProceedTemplate([][]string)  string
    GetData()([]Line,error)
    GetType()(typer int)
}


type Line struct {
    data_string_slice []string
    data_indexes      [][]int
    delim_indexes     [][]int
    data              [][]string
    template          string
}

type Section struct {
    lines             []Line
    template          string
    typer             int
}


func(l *Line)GetData()(lines []Line,err error){
    if (l!=nil){
        lines = make([]Line,   0)
        lines = append(lines, *l)
        return lines, nil
    } else {
        return lines, nilResultError
    }
}

func(s *Section)GetData()([]Line,err error){
    if (s.lines!=nil){
        lines = make([]Line,0)
        lines = append(lines, *l)
        return lines, nil
    } else {
        return lines, nilResultError
    }
}
