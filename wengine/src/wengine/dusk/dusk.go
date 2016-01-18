package dusk

import "wengine/core/utah"


type Database interface {

    Connect()(error)
    Close()()
    CreateUser(user *utah.User)(err error)
    //GetUsers( []utah.User,error)
    //GetGroups([]utah.Group,error)
    //CreateUser(username string, password string)(id string ,err error)
    //GetUser(id string)(user utah.User)
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
