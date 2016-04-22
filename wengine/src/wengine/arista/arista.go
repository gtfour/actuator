package arista
import "errors"

var base_word = "arista"

func errwrap(in string)(string) {
    return base_word+":"+in
}

var access_allowed = errors.New(errwrap("access_allowed"))
var access_denied  = errors.New(errwrap("access_denied"))


type Stance interface {
    SetName(string)(error)
    GetName(string)(error)
    SetType(string)(error)
    GetType()(string,error)
    SetId(string)(error)
    GetId()(string,error)
    GetGroupIds()([]string,error)
    AddGroupId(string)(error)
    RemoveGroupId(string)(error)
}

type Member interface {
    Stance
}

type group interface {
    GetMember()(*Member,error)
    AddMember(Member)(error)
    RemoveMember(Member)(error)
    GetMembers()([]Member)
}

type Group interface {
    Stance
    group
}


type resource interface {
    GrantAccessToMember(string)(error)
    RemoveAccessForMember(string)(error)
    GrantAccessToGroup(string)(error)
    RemoveAccessForGroup(string)(error)
    CheckPermissionForMember(string)(error)
    CheckPermissionForGroup(string)(error)
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


func GetMember(string)(m Member) {
    return m
}

func GetResource(string)(r Resource){
    return r
}
