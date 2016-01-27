package dusk

import "wengine/core/utah"
//import "wengine/core/dashboard"


type Database interface {

    Connect()(error)
    Close()()
    CreateUser(*utah.User)(user_id string,err error)
    RemoveUsersById(id ...string)                   (err error)
    //GetUsers( []utah.User,error)
    //GetGroups([]utah.Group,error)
    //CreateUser(username string, password string)(id string ,err error)
    GetUserById(id string)(utah.User,error)
    GetUser(map[string]interface{})                 (utah.User,error)
    //GetUserDashboards(user_id string)([]dashboard.Dashboard,error)
    AttachDashboardToUser(user_id,dash_id string)(error)

    //CreateDashboard(*dashboard.Dashboard)(error)
    //RemoveDashboardById(dash_id string)(error)


    CreateToken(userid string)                      (token string,err error)
    TokenIsExist(userid string, token_id string)    (bool)
    RemoveToken(user_id,token_id string)(error)
    GetAnUserToken(string)(string,error)
    UserPasswordIsCorrect(username ,password string)( user_id string, token_id string,exists bool)
    //GetUserToken(userid string)(token string,error)
    //RemoveUserToken(userid string)(token string,error)
    //RemoveUsers( map[string]interface{} )
    //
    // dashboard
    //CreateDashboard(*dashboard.Dashboard)(dashboard_id string,err error )
    //GetDashboardById(dashboard_id string)(dashboard.Dashboard)

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
