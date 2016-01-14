package dusk

import "wengine/core/utah"

//type DBConnection struct {}
//func CreateDBConnection (dbtype string )( d DBConnection ) { return }



type Database interface {

    GetUsers( []utah.User,error)
    GetGroups([]utah.Group,error)
    CreateUser(username string, password string)(id string ,err error)
    GetUser(id string)(user utah.User)
}

func GetUsers ( d Database )() {


}

func OpenDatabase ( dbtype string , username string, password string, host string, dbname string  ) ( d Database ) {

    switch {
        case db == "mongo":

        case db == "postgres":


    }
    return nil


}
