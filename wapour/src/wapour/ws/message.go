package ws
import "encoding/json"

type Message struct {
    DataType   string          `json:"datatype"`
    Data       json.RawMessage `json:"data"`
    CallbackId int             `json:"callback_id"`
    SessionId  string          `json:"session_id"`
}

type MessageSwitchDashboard struct {
    // "dashboard_group_id":dashboard_group_id, "dashboard_id":dashboard_id
    DashboardGroupId string `json:"dashboardgroupid"`
    DashboardId      string `json:"dashboardid"`
}

//"datatype":"message_ws_open","session_id":session_id

type MessageWsState struct {
    // "dashboard_group_id":dashboard_group_id, "dashboard_id":dashboard_id
//    DashboardGroupId string `json:"dashboardgroupid"`
    State string `json:"state"`
}


type MessageChat struct {
    Author string `json:"author"`
    Body   string `json:"body"`
}

func (self *MessageChat) String()string {
    return self.Author + " says "+self.Body
}
