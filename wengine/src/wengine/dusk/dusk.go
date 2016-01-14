package dusk

import "wengine/core/utah"

//type DBConnection struct {}
//func CreateDBConnection (dbtype string )( d DBConnection ) { return }

type Database interface {

    GetUsers( []utah.User,error)
    GetGroups([]utah.Group,error)

}

func GetUsers ( d Database )() {


}
