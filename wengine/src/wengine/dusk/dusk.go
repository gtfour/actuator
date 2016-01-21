package dusk

import "wengine/core/utah"


type Database interface {

    Connect()(error)
    Close()()
    CreateUser(*utah.User)(error)
    //GetUsers( []utah.User,error)
    //GetGroups([]utah.Group,error)
    //CreateUser(username string, password string)(id string ,err error)
    GetUserById(id string)(utah.User,error)
    GetUser(map[string]interface{})(utah.User,error)
    RemoveUsersById(id ...string)(err error)
    CreateToken(userid string)(token string,err error)
    //GetUserToken(userid string)(token string,error)
    //RemoveUserToken(userid string)(token string,error)
    //RemoveUsers(map[string]interface{})
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
