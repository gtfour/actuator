package aristo


import "fmt"
import "wengine/dusk"
import "wengine/core/common"
import "wengine/core/types/db_types"

type Member struct {
    Type  string
    Id    string
    Name  string
}


type Group struct {
    Name    string
    Id      string
    Members []Member
}




func CreateNewGroup()(g *Group,err error) {
   //group_prop := s.GetProp

   new_group         := make(map[string]interface{},0)
   new_group["id"],_ =  common.GenId()
   new_query         := dusk.Query{Table:GROUPS_T, Type:db_types.CREATE_NEW, QueryBody:new_group}
   /*result*/_,err   = database.RunQuery(new_query)

   return g,err

}


func GetGroup(prop map[string]interface{},query_type ...int)(gs map[string]interface{},err error){

    //key`_body        := make(map[string]interface{},0)
    //key_body["id"]  =  id
    if prop != nil {
        selected_query_type:=db_types.GET
        if len(query_type) == 1 {
            prov_query_type:=query_type[0]
            if prov_query_type == db_types.GET || prov_query_type == db_types.GET_ALL {
                selected_query_type = prov_query_type
            }
        }
        new_query         := dusk.Query{Table:GROUPS_T, Type:selected_query_type, KeyBody:prop}
        result_slice, err := database.RunQuery(new_query)

        fmt.Printf("Result:\n%v\n",result_slice)
        gs = nil
        if gs == nil {
            return nil, group_list_is_empty
        } else {
            return gs,err
        }
    }else{
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

func LoadMember(prop map[string]interface{})(members []Member,err error){
    return members,err
}

func AddMember()(members []Member,err error){
    return members,err
}

