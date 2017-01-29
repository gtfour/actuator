package aristo

type Signa struct{
    Id       string
    GroupId  string
    Type     int
}

type Club struct {
    // sheep is a member
    Sheep      string
    SheepType  int
    //
    //
    //
    Herd       string
    HerdType    int
    //
}

type Cord struct {
    //
    //
    //
    signaReqId         string
    signaReqType       int
    //signaReqGroupId    string
    //
    //
    signaTgtId         string
    signaTgtType       int
    //signaTgtGroupId    string
    //
    //
    //
}

func CheckAccess(signaReq *Signa, signaTgt *Signa)(state int){
    return SirAccessAllowed
}
