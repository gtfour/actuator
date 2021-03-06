package dusk

import "wengine/settings"
import "wengine/activa"
import "wengine/core/utah"
import "wengine/core/dashboard"

var DATABASE_INSTANCE = OpenDatabase(settings.PrimaryDatabase,
                                     settings.DBusername,
                                     settings.DBpassword,
                                     settings.DBhost,
                                     settings.DBname)


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
    ActivaMethods
    AristoMethods
    //
    QueryMethods
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

    // DashboardGroupMethods
    CreateDashboard(dashboard *dashboard.Dashboard)(dashboard_id string, err error)
    DashboardExists(dashboard_id string)(bool)
    AttachDashboardToUser(user_id,dash_id string)(error)
    GetMyDashboardList(user_id,token_id string)(dashboard_list dashboard.DashboardList, err error)
    //GetDashboardData(dashboard_group_id string,dashboard_id string)()
    //UpdateDashboardData(dashboard_group_id string,dashboard_id string, dashboard_data )

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

type AristoMethods interface {
    CheckAccess(initiator_type,initiator_id,target_type,target_id string)(error)
    GrantAccess(prop map[string]string)(error,int)
    RemoveAccess(prop map[string]string)(error,int)
}

type ActivaMethods interface {

    WriteMotion (*activa.Motion)(error)
    //GetMotion   (...activa.Key)(error)
    //UpdateMotion(*activa.Motion)(error)
    //RemoveMotion(...activa.Key)(error)

}


type SolisMethods interface {
}

type QueryMethods interface {
    RunQuery(Query)(result_slice *[]map[string]interface{}, err error)
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
