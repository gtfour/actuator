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

type Motion struct {


}

func CreateMotion () {

}

func (m *Motion) Test ()() {

}

func Update()() {


}



