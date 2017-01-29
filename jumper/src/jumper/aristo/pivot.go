package aristo

type Signa struct{
    Sid         string
    PrimaryGid  string
    Stype       int
}

type Membership struct {
    //
    Mid    string
    Mtype  int
    //
    //
    //
    Gid    string
    Gtype  int
    //
}

type Cord struct {
    //
    //
    // cord show's access rights(relation) between an requestor(Req) and resource(Tgt)
    //
    //
    signaReqId          string
    signaReqType        int
    // signaReqGroupId  string
    //
    //
    signaTgtId          string
    signaTgtType        int
    // signaTgtGroupId  string
    //
    //
    //
}

func CheckAccess(signaReq *Signa, signaTgt *Signa)(state int){
    return SirAccessAllowed
}
