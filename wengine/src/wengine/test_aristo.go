package main

import "fmt"
import "wengine/aristo"
import "wengine/core/types/db_types"

//var database dusk.Database = dusk.DATABASE_INSTANCE

func main() {
    //database.CheckAccess("11","12","41","42")
    key_query       := make(map[string]interface{})
    //key_query["id"] =  "E3B41931-53DE-6912-79B7-415521CA8FDE"
    //_,_=aristo.CreateNewGroup()
    g:=aristo.Group{Id:""}
    groups,err:=aristo.GetGroup(key_query,db_types.GET_ALL)
    fmt.Printf("\n%v\n%v",groups,err)

}
