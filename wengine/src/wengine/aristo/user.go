package aristo

type User struct {
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

func ( u *User ) SetName ( name string )(error){
    u.name = name
    return nil
}

func ( u *User ) GetName ()( string, error){
    return u.name,nil
}

func GetUser(string)(m Member) {
    return m
}

func CreateUser(string)(m Member) {
    return m
}


func (m *User)EditUser(prop ...map[string]string)() {
}



