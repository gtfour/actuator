package chase

import "client_side/actuator"



type Target struct {

    Path string
    OldMarker string
    Marker string
    Modified bool
    EventGroup string
    EventType string
    IsDir bool
    EventChan chan
    InfoChan chan

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
                go target.ChasingiFile()


            }

        }
    }else {

      target:=&Target{}
      target.Path=targets[id]
      target.OldSum=string(file_struct.Sum)
       go target.ChasingFile()


    }
    }

    return nil

}

func Stop()(err error) {


    return nil
}


func (tgt *Target) ChasingFile (){


}

func (tgt *Target) ChasingDir (message chan string){

    
    for {

        tgt.Marker=actuator.Get_mtime(tgt.Path)
        

    }


}


func (tgt *Target) Reporting (){

    //UpdateConfFile()
    //SendPostRequest()

}
