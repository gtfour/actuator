package chase

import "client_side/actuator"



type Target struct {

    Path string
    OldSum string
    Sum string
    Modified bool
    EventGroup string
    EventType string

}

func Start (targets []string)(err error){

    for id :=range targets {

    file_struct,err:= actuator.Get_md5_file(targets[id])
    if err!=nil {

        dir_struct,err:=actuator.Get_md5_dir(targets[id])
        if err==nil {

            for file_id :=range dir_struct.Files{

                target:=&Target{}
                target.Path=dir_struct.Files[file_id].Path
                target.OldSum=string(dir_struct.Files[file_id].Sum)
                go target.Chasing()


            }

        }
    }else {

      target:=&Target{}
      target.Path=targets[id]
      target.OldSum=string(file_struct.Sum)
       go target.Chasing()


    }
    }

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
