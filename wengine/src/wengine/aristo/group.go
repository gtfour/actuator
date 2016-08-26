package aristo


import "wengine/dusk"
import "wengine/core/common"
import "wengine/core/types/db_types"

type member struct {


}

type Group struct {
    Name    string
    //Members []string
}




func CreateNewGroup()(g *Group,err error) {
   //group_prop := s.GetProp

   new_group         := make(map[string]interface{},0)
   new_group["id"],_ =  common.GenId()
   new_query         := dusk.Query{Table:GROUPS_T, Type:db_types.CREATE_NEW, QueryBody:new_group}
   /*result*/_,err   = database.RunQuery(new_query)

   return g,err

}


func GetGroup(id string)(g *Group,err error){

    key_body        := make(map[string]interface{},0)
    key_body["id"]  =  id
    new_query       := dusk.Query{Table:GROUPS_T, Type:db_types.GET, KeyBody:key_body}
    /*result*/_,err = database.RunQuery(new_query)

    return g,err

}

func GetGroupProp(id string)(props map[string]interface{},err error) {
    return props,err
}



