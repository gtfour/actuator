package activa
type Motion struct {
    /*Id           string
    InitTime     string
    EndTime      string
    UserName     string
    GroupName    string*/
    Id           string
    InitTime     string
    StartTime    string
    EndTime      string
    UserName     string
    GroupName    string
    SourceType   string
    SourcePath   string
    ActionType   string
    ActionName   string
    TaskState    int
    //Task
}
type Task struct {

    Type       string
    SourceType string
    SourcePath string

}



