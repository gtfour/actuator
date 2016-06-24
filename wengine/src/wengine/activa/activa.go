package activa
import "wengine/core/utah"
import "wengine/core/dashboard"

type AbMotion interface {
    GetId           (string)
    GetInitTime     (string)
    GetSuccessTime  (string)
    GetInput        (interface{})
    GetOutput       (interface{})
    GetScope        (interface{})
    GetOwner        (utah.User)
    GetDashboards   ([]dashboard.Dashboard)
}

type Key struct {
    Name  string
    Value string
}

type Motion struct {
    Id           string //`bson:"id"`
    InitTime     string //`bson:"init_time"`
    EndTime      string //`bson:"end_time"`
    UserName     string //`bson:"user_name"`
    GroupName    string //`bson:"group_name"`
}

func CreateMotion ()() {

}

func (m *Motion) Test ()() {

}

func Update()() {


}



