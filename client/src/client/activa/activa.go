package activa

type Motion struct {
    Id           string
    InitTime     string
    EndTime  string
    Task
}

type Task struct {

    Type       string
    SourceType string
    SourcePath string



}



