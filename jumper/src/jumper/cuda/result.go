package cuda

type Result interface {
    ProceedTemplate([][]string)  string
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
