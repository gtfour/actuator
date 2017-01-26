package aristo

type User struct {
    id           string
    name         string
    utype        string
    email        string
}

type HumanUser struct {
//    SetName(string)(error)
//    GetName(string)(error)
//    SetType(string)(error)
//    GetType()(string,error)
//    SetId(string)(error)
//    GetId()(string,error)
//    GetGroupIds()([]string,error)
//    AddGroupId(string)(error)
//    RemoveGroupId(string)(error)
    //name string
    User
    second_name  string
}

type SystemUser struct {
    User
}

func ( u *User )SetName( name string )(error){
    u.name = name
    return nil
}

func ( u *User )GetName()( string, error){
    return u.name,nil
}

func GetUser(string)(u User) {
    return u
}

func CreateUser(string)(u User) {
    return u
}


func (u *User)EditUser(prop map[string]interface{})() {
}





func(u *User)UserToMember()(m Member, err error) {
    m.Id   = u.id
    m.Name = u.name
    m.Type = u.utype
    return m,err
}

