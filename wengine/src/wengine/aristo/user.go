package aristo

type Human struct {
//    SetName(string)(error)
//    GetName(string)(error)
//    SetType(string)(error)
//    GetType()(string,error)
//    SetId(string)(error)
//    GetId()(string,error)
//    GetGroupIds()([]string,error)
//    AddGroupId(string)(error)
//    RemoveGroupId(string)(error)
    name string
}

func ( h *Human )SetName( name string )(error){
    h.name = name
    return nil
}

func ( h *Human )GetName()( string, error){
    return h.name,nil
}

func GetUser(string)(u User) {
    return u
}

func CreateUser(string)(u User) {
    return u
}


func (h *Human)EditUser(prop ...map[string]string)() {
}



