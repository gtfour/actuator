package aristo

type Group struct {


}


func CreateNewGroup()(g *Group,err error) {
   //group_prop := s.GetProp
   new_query:=Query{Table:GROUPS_T}
   new_query.Run()
   return g,err
}

func GetGroup(id string)(g *Group,err error){
    return g,err
}

func GetGroupProp(id string)(props map[string]interface{},err error) {
    return props,err
}

