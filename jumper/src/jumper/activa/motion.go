package activa
//import "fmt"
import "encoding/json"
import "jumper/common/gen"

/*
type AbMotion interface {
    GetId           (string)
    GetInitTime     (string)
    GetSuccessTime  (string)
    GetInput        (interface{})
    GetOutput       (interface{})
    GetScope        (interface{})
    //GetOwner        (utah.User)
    //GetDashboards   ([]dashboard.Dashboard)
}

type Key struct {
    Name  string
    Value string
}
*/

type Motion struct {
    Id               string // `bson:"id"`
    InitTime         string // `bson:"init_time"`
    StartTime        string // `bson:"start_time"`
    EndTime          string // `bson:"end_time"`
    UserName         string // `bson:"user_name"`
    GroupName        string // `bson:"group_name"`
    SourceType       string // `bson:"source_type"`
    SourcePath       string // `bson:"source_path"`
    ActionType       string // `bson:"action_type"`
    ActionName       string // `bson:"action_name"`
    TaskState        int    // `bson:"task_state"`
    Type             int
    MotionData       json.RawMessage `json:"data"`
    MotionLastUpdate string
}

func (m *Motion)GetRaw()([]byte, error) {
    raw,err:=json.Marshal(m)
    return raw,err
}

func CreateNewMotion ()(m Motion) {

    time_now    := gen.GetTime()
    m.Id        =  time_now
    m.InitTime  =  time_now
    m.TaskState =  TASK_STATE_new
    return m


}

func (m *Motion) Test ()() {

}

func Update()() {


}

