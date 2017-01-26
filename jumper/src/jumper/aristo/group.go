package aristo


import "fmt"
import "wengine/dusk"
import "wengine/core/common"
import "wengine/core/types/db"

type Member struct {
    Id    string
    Name  string
    Type  string
}


type Group struct {
    Id      string
    Name    string
    Type    string
    Members []Member
}



func CreateNewMember(name string, mtype string)(m Member){
    member_type:="general"
    if mtype != "" { member_type=mtype }
    new_id,_   := common.GenId()
    m.Id = new_id
    m.Name = name
    m.Type = member_type
    return m
}

func (m *Member)Write()(err error){

    member_map,err:=m.Check()
    if err == nil {
        new_query:=dusk.Query{Table:MEMBERS_T, Type:db.CREATE_NEW, QueryBody:member_map}
        _,err   = database.RunQuery(new_query)
        return err
    } else { return err }
}

func CreateNewGroup(name string, gtype string)(g Group) {
   // default type: general
   //group_prop := s.GetProp
    group_type :="general"
    if gtype   != "" { group_type=gtype }
    new_id,_   :=  common.GenId()
    g.Members  =make([]Member,0)
    g.Id       = new_id
    g.Type     = group_type
    g.Name     = name
    return g

}

func (g *Group)Write()(err error){
    group_map,err:=g.Check()
    if err == nil {
        new_query         := dusk.Query{Table:GROUPS_T, Type:db.CREATE_NEW, QueryBody:group_map}
        _,err   = database.RunQuery(new_query)
        return err
    } else { return err }


}


func GetGroup(prop map[string]interface{},query_type ...int)(gs map[string]interface{},err error){

    //key`_body        := make(map[string]interface{},0)
    //key_body["id"]  =  id
    if prop != nil {
        selected_query_type:=db.GET
        if len(query_type) == 1 {
            prov_query_type:=query_type[0]
            if prov_query_type == db.GET || prov_query_type == db.GET_ALL {
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
    // old_id and new_id should be equal
    old_id, ok_old := prop["id"]
    new_id, ok_new := new_prop["id"]
    if ok_old == false {  //|| ok_new == false  {
        return id_isnot_specified
    }
    if ok_new == true && old_id!= new_id {
        return id_change_is_not_allowed
    }
    new_query         := dusk.Query{Table:GROUPS_T, Type:db.EDIT, KeyBody:prop ,QueryBody:new_prop}
    _,err   = database.RunQuery(new_query)
    return err
}

func GetGroupProp(id string)(props map[string]interface{},err error) {
    return props,err
}

func LoadMember(prop map[string]interface{})(members []Member,err error){
    return members,err
}

func (g *Group)AddMember(m Member)(err error){
    group,err_group       := g.Check()
    member,err_member     := m.Check()
    query_body            := make(map[string]interface{}, 0)
    query_body["members"] = member
    if err_group == nil || err_member == nil {
        new_query         := dusk.Query{Table:GROUPS_T, Type:db.INSERT_ITEM, KeyBody:group ,QueryBody:query_body}
        _,err             = database.RunQuery(new_query)
    }
    return err
}

func (g *Group)Remove()(err error){

    group,err := g.Check()
    if err == nil {
        new_query         := dusk.Query{Table:GROUPS_T, Type:db.REMOVE, KeyBody:group}
        _,err             = database.RunQuery(new_query)
        return err
    } else { return err }
}

func (g *Group)RemoveMember(member Member)(err error){

    return err
}


func(m *Member)GetMemberGroups()(gs []Group,err error) {
    return gs,err
}

func (g *Group)Check()(group map[string]interface{}, err error){
    if g.Id == "" || g.Type == "" {
        return nil, group_invalid
    } else {
        group = make(map[string]interface{},0)
        group["id"]   = g.Id
        group["name"] = g.Name
        group["type"] = g.Type

        if len(g.Members) > 0 { group["members"] = make([]map[string]interface{},0) }

        member_list:=make([]map[string]interface{},0)

        for i:= range g.Members {
            member         := g.Members[i]
            member_map,err := member.Check()
            if err == nil {
                member_list = append(member_list, member_map)
            }
        }

        if len(member_list) > 0 { group["members"]=member_list }

        return group, nil
    }
}

func (m *Member)Check()(member map[string]interface{},err error){
    if  m.Id == "" || m.Type == "" {
        return nil, member_invalid

    } else {
        member = make(map[string]interface{},0)
        member["id"]   = m.Id
        member["name"] = m.Name
        member["type"] = m.Type
        return member, nil
    }
}

func MakeMember(member_prop map[string]interface{})(m *Member,err error){
    return m,err
}

func MakeGroup(group_prop map[string]interface{})(g *Group,err error){
    return g,err
}




