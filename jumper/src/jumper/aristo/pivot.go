package aristo

type Signa struct{
    Id       string
    GroupId  string
    Type     int
}

func CheckAccess(signaReq *Signa, signaTgt *Signa)(state int){
    return SirAccessAllowed
}
