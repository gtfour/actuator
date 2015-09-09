package chase



type Target struct {

    Path string
    OldSum string
    Sum string
    Modified bool

}

func Start (targets []string)(err error){

    return nil

}

func Stop()(err error) {


    return nil
}


func (tgt *Target) Chasing (){


}

func (tgt *Target) Reporting (){

    //UpdateConfFile()
    //SendPostRequest()

}
