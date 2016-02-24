package dusk

import "wengine/core/utah"
import "wengine/core/dashboard"


type Database interface {

    Connect()(error)
    Close()()
    //RemoveUsersById(id ...string)                   (err error)
    //GetUsers( []utah.User,error)
    //GetGroups([]utah.Group,error)
    //CreateUser(username string, password string)(id string ,err error)
    //GetUserById(id string)(utah.User,error)
    //GetUser(map[string]interface{})                 (utah.User,error)
    //GetUserDashboards(user_id string)([]dashboard.Dashboard,error)
    //CreateDashboard(dashboard *dashboard.Dashboard)(dashboard_id string, err error)
    // DashboardGroupMethods
    //DashboardExists(dashboard_id string)(bool)
    //AttachDashboardToUser(user_id,dash_id string)(error)
    //GetMyDashboardList(user_id,token_id string)(dashboard_list dashboard.DashboardList, err error)
    UserMethods
    //CreateDashboard(*dashboard.Dashboard)(error)
    //RemoveDashboardById(dash_id string)(error)
    DashboardMethods
    DashboardGroupMethods
    TokenMethods


    //GetUserToken(userid string)(token string,error)
    //RemoveUserToken(userid string)(token string,error)
    //RemoveUsers( map[string]interface{} )
    //
    // dashboard
    //CreateDashboard(*dashboard.Dashboard)(dashboard_id string,err error )
    //GetDashboardById(dashboard_id string)(dashboard.Dashboard)

}

type UserMethods interface {

    CreateUser(*utah.User)(user_id string,err error)
    RemoveUsersById(id ...string)                   (err error)
    //GetUsers( []utah.User,error)
    //GetGroups([]utah.Group,error)
    //CreateUser(username string, password string)(id string ,err error)
    GetUserById(id string)(utah.User,error)
    GetUser(map[string]interface{})                 (utah.User,error)
    //GetUserDashboards(user_id string)([]dashboard.Dashboard,error)
    GetAllUsers()([]utah.User,error)

}

type DashboardMethods interface {

    CreateDashboard(dashboard *dashboard.Dashboard)(dashboard_id string, err error)
    // DashboardGroupMethods
    DashboardExists(dashboard_id string)(bool)
    AttachDashboardToUser(user_id,dash_id string)(error)
    GetMyDashboardList(user_id,token_id string)(dashboard_list dashboard.DashboardList, err error)

}

type DashboardGroupMethods interface {
    CreateDashboardGroup     (dashboard *dashboard.DashboardGroup)          (dashboardgroup_id string, err error)
    RemoveDashboardGroup     (dashboardgroup_id string)                     (error)
    AddDashboardToGroup      (dashboardgroup_id string,dashboard_id string) (error)
    AttachDashboardGroupToUser (user_id,dgroup_id string)(error)
    RemoveDashboardFromGroup (dashboardgroup_id string,dashboard_id string) (error)
    GetMyDashboardGroupList  (user_id,token_id string)                      (dashboard.DashboardGroupList,error)
}

type TokenMethods interface {
    CreateToken           (userid string)                   (token string,err error)
    TokenExists           (userid string, token_id string)  (bool)
    RemoveToken           (user_id,token_id string)         (error)
    GetAnUserToken        (string)                          (string,error)
    UserPasswordIsCorrect (username ,password string)       ( user_id string, token_id string,exists bool)
}


func OpenDatabase ( dbtype, username, password, host, dbname  string) ( d Database ) {

    switch {
        case dbtype == "mongo":
            d=&MongoDb{username:username,
                      password:password,
                      host:host,
                      dbname:dbname}
            err:=d.Connect()
            if err == nil {
                return d
            }
        case dbtype == "postgres":
    }
    return nil
}
