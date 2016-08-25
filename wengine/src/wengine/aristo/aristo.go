package aristo
import "errors"

var base_word = "aristo"

func errwrap(in string)(string) {
    return base_word+":"+in
}

var access_allowed = errors.New(errwrap("access_allowed"))
var access_denied  = errors.New(errwrap("access_denied"))


type Stance interface {
    //SetName(string)(error)
    //GetName(string)(error)
    //SetType(string)(error)
    //GetType()(string,error)
    //SetId(string)(error)
    //GetId()(string,error)
    //GetGroupIds()([]string,error)
    //AddGroupId(string)(error)
    //RemoveGroupId(string)(error)
    SetProp(map[string]string)(error)
    GetProp()(map[string]string,error)
}

type User interface {
    Stance
}

type group interface {
    //Stance
    //GetUser()(*User,error)
    //AddUser(User)(error)
    //RemoveUser(User)(error)
    //GetUsers()([]User)
}

type Group interface {
    Stance
    //group
}


type resource interface {
    GrantAccessToUser(string)(error)
    RemoveAccessForUser(string)(error)
    GrantAccessToGroup(string)(error)
    RemoveAccessForGroup(string)(error)
    CheckPermissionForUser(string)(error)
    CheckPermissionForGroup(string)(error)
    //
    GetRaw()([]byte, error)
}

type Action interface {
    Stance
    SetResource(string)(error)
    GetResource()(string,error)
}


type Resource interface {
    Stance
    resource
}

type ResourceGroup interface {
    Stance
    resource
    group
}


//func GetUser(string)(m User) {
//    return m
//}

func GetResource(string)(r Resource){
    return r
}
