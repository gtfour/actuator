package activa
import "fmt"
import "encoding/json"
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
    Type         int
    // SourceType   string
    // SourcePath   string
    // ActionType   string
    // ActionName   string
    TaskState    int
    //Task
    MotionData   json.RawMessage `json:"data"`
}
type Task struct {

    Type       string
    SourceType string
    SourcePath string

}

func Handle( motions chan *Motion )() {
    for {
        select {
            case motion:=<-motions:
                //cross.WriteMotion(motion)
                fmt.Printf("<<New motion:\n%v\n>>", motion)

            }
        }
}
