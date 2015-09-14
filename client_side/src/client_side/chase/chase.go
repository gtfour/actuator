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
    InfoIn <-chan bool
    InfoOut chan<- string

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
                target.OldMarker=string(dir_struct.Files[file_id].Sum)
                go target.ChasingFile()


            }

        }
    }else {

      target:=&Target{}
      target.Path=targets[id]
      target.OldMarker=string(file_struct.Sum)
       go target.ChasingFile()


    }
    }

    return nil

}

func Stop()(err error) {


    return nil
}


func (tgt *Target) ChasingFile() (err error){



    for {

        ask_path:= <-tgt.InfoIn

        if(ask_path) { tgt.InfoOut <- tgt.Path }

        if file,err:=actuator.Get_md5_file(tgt.Path);err!=nil { tgt.Marker=string(file.Sum) } else { return err }

        if (tgt.Marker!=tgt.OldMarker){ tgt.Reporting() }

        tgt.OldMarker=tgt.Marker


    }
    return nil


}

func (tgt *Target) ChasingDir()(err error){

    
    for {

        tgt.Marker=actuator.Get_mtime(tgt.Path)

        if (tgt.Marker!=tgt.OldMarker){ tgt.Reporting() }

        tgt.OldMarker=tgt.Marker
        

    }


}


func (tgt *Target) Reporting (){

    //UpdateConfFile()
    //SendPostRequest()

}
