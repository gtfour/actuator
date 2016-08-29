package main

import "fmt"
import "wengine/aristo"
import "wengine/core/types/db_types"

//var database dusk.Database = dusk.DATABASE_INSTANCE

func main() {
    //database.CheckAccess("11","12","41","42")
    key_query:=make(map[string]interface{})
    //_,_=aristo.CreateNewGroup()
    groups,err:=aristo.GetGroup(key_query,db_types.GET_ALL)
    fmt.Printf("\n%v\n%v",groups,err)

}
