package webclient
import "fmt"
import "net/http"
import "encoding/json"
import "io/ioutil"
import "wapour/settings"
//import "wapour/salvo"
import "wapour/api/wengine"


var USER_DASHBOARD_LIST_URL string = "/rest/user/get-my-dashboards"
var USER_GET_DASHBOARD_DATA string = "/rest/user/get-dashboard-data"

type DashboardListResult struct {
    Status             string                     `json:"status"`
    DashboardList      wengine.DashboardList      `json:"dashboard_list"`
    DashboardGroupList wengine.DashboardGroupList `json:"dashboard_group_list"`
}


func GetUserDashboards( user_id string, token_id string )(dashboards DashboardListResult){
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

func GetDashboardData(user_id string , token_id string , dashboard_id string )([]byte,error){
    url  := settings.RESTAPI_URL + USER_GET_DASHBOARD_DATA+"/"+dashboard_id
    client := &http.Client{}
    req,_ := http.NewRequest("GET", url, nil)
    SetAuthCookie(req,user_id,token_id)
    resp, _ := client.Do(req)
    defer resp.Body.Close()
    //decoder := json.NewDecoder(resp.Body)
    //_ = decoder.Decode(&dashboards)
    //return dashboards
    data,err:=ioutil.ReadAll(resp.Body)
    return data,err
}

