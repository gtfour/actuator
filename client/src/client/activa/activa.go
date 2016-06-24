package activa
type Motion struct {
    Id           string
    InitTime     string
    EndTime      string
    UserName     string
    GroupName    string
    Task
}
type Task struct {

    Type       string
    SourceType string
    SourcePath string

}



