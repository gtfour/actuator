package webclient
import "fmt"
import "net/http"
import "encoding/json"
import "wapour/api/wengine"
import "wapour/settings"


var USER_DASHBOARD_LIST_URL string = "/rest/user/get-my-dashboards"

type DashboardListResult struct {
    Status             string                 `json:"status"`
    DashboardList      wengine.DashboardList  `json:"dashboard_list"`
    DashboardGroupList []string               `json:""dashboard_group_list""`
}


func GetUserDashboards(token_id string,user_id string, wrappers *[]*WengineWrapper)(dashboards DashboardListResult){

    url  := settings.RESTAPI_URL + USER_DASHBOARD_LIST_URL
    client := &http.Client{}
    req,_ := http.NewRequest("GET", url, nil)
    SetAuthCookie(req,user_id,token_id)
    fmt.Printf("\nRequest : %v\n",req)
    resp, _ := client.Do(req)
    defer resp.Body.Close()
    decoder := json.NewDecoder(resp.Body)
    _ = decoder.Decode(&dashboards)
    return dashboards

}
