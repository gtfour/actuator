package aristo


import "wengine/dusk"
import "wengine/core/common"
import "wengine/core/types/db_types"

type member struct {


}

type Group struct {
    Name    string
    Members []string
}




func CreateNewGroup()(g *Group,err error) {
   //group_prop := s.GetProp

   new_group         := make(map[string]interface{},0)
   new_group["id"],_ =  common.GenId()
   new_query         := dusk.Query{Table:GROUPS_T, Type:db_types.CREATE_NEW, QueryBody:new_group}
   /*result*/_,err   = database.RunQuery(new_query)

   return g,err

}


func GetGroup(prop map[string]interface{})(gs map[string]interface{},err error){

    //key`_body        := make(map[string]interface{},0)
    //key_body["id"]  =  id
    if prop != nil {
        new_query       := dusk.Query{Table:GROUPS_T, Type:db_types.GET, KeyBody:prop}
        gs,err          = database.RunQuery(new_query)
        if gs == nil {
            return nil, group_list_is_empty
        } else {
            return gs,err
        }
    } else {
        return nil, prop_is_empty

    }

}

func EditGroup(prop map[string]interface{}, new_prop map[string]interface{})(err error){

    //key_body := make(map[string]interface{},0)
    _, ok_old := prop["id"]
    _, ok_new := new_prop["id"]
    if ok_old == false || ok_new == false  {
        return id_isnot_specified
    }
    new_query         := dusk.Query{Table:GROUPS_T, Type:db_types.EDIT, KeyBody:prop ,QueryBody:new_prop}
    _,err   = database.RunQuery(new_query)
    return err
}

func GetGroupProp(id string)(props map[string]interface{},err error) {
    return props,err
}



