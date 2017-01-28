package aristo

type Signa struct{
    Id       string
    GroupId  string
    Type     int
}

type Cord struct {
    //
    //
    //
    signaReqId         string
    signaReqGroupId    string
    //
    //
    signaTgtId         string
    signaTgtGroupId    string
    //
    //
    //
}

func CheckAccess(signaReq *Signa, signaTgt *Signa)(state int){
    return SirAccessAllowed
}
