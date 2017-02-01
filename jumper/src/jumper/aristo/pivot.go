package aristo

type Signa struct{
    // atomic instance 
    Sid         string
    PrimaryGid  string
    Stype       int
}

type Membership struct {
    Mid    string  // member id
    Mtype  int     // member type
    Gid    string  // group id
    Gtype  int     // group type
}

type Cord struct {
    //cord show's access rights(relation) between an requestor(Req) and resource(Tgt)
    signaReqId         string
    signaReqType       int
    //signaReqGroupId  string
    signaTgtId         string
    signaTgtType       int
    //signaTgtGroupId  string
    cordType           int
}

func CheckAccess(signaReq *Signa, signaTgt *Signa)(state int){
    return SirAccessAllowed
}
//func GetDefaultStrategy
