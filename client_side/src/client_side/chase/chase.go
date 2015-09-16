//package chase
package main

import "client_side/actuator"
import "os"
import "fmt"
import "time"


type Message struct {

    Path string
    IsChased bool
    KillSelf bool

}

type Target struct {

    Path string
    OldMarker string
    Marker string
    Modified bool
    EventGroup string
    EventType string
    IsDir bool
    InfoIn chan bool
    InfoOut chan string
    MessageChannel chan string
    

}

type TargetDir struct {

    Path string
    OldMarker string
    Marker string
    InfoIn []chan bool
    InfoOut []chan string
    MessageChannel chan string


}


func Start (targets []string, message_channel chan string)(err error){


    //request_channel:=make(chan bool)

    //response_channel:=make(chan string)

    for id :=range targets {

    file_struct,err:= actuator.Get_md5_file(targets[id])

    if err!=nil {

        dir_struct,err:=actuator.Get_md5_dir(targets[id])

        subdirs:=make(map[string]TargetDir)

        for subname:=range dir_struct.SubDirs  {

            tgt_dir:=TargetDir{}
            tgt_dir.MessageChannel=message_channel

            subdirs[dir_struct.SubDirs[subname]]=tgt_dir
            go tgt_dir.ChasingDir()



        }

        if err==nil {

            for file_id :=range dir_struct.Files{

                file_struct:=dir_struct.Files[file_id]

                target:=&Target{}
                target.Path=file_struct.Path
                target.OldMarker=string(file_struct.Sum)
                target.MessageChannel=message_channel

                if subdir, ok := subdirs[file_struct.Dir]; ok {

                    subdir.InfoIn=append(subdir.InfoIn,target.InfoIn)
                    subdir.InfoOut=append(subdir.InfoOut,target.InfoOut)


                }//else {
                //    target.InfoIn = request_channel
                //    target.InfoOut = response_channel
                //}
                go target.ChasingFile()


            }

        }
    }else {

      target:=&Target{}
      target.Path=targets[id]
      target.OldMarker=string(file_struct.Sum)
      //target.InfoIn = request_channel
      //target.InfoOut = response_channel
      target.MessageChannel=message_channel
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

        
        select {
            case ask_path:= <-tgt.InfoIn:

                if(ask_path) { tgt.InfoOut <- tgt.Path }
            default:

                if file,err:=actuator.Get_md5_file(tgt.Path);err!=nil { tgt.Marker=string(file.Sum) } else { return err }

                if (tgt.Marker!=tgt.OldMarker){ tgt.Reporting() }

                tgt.OldMarker=tgt.Marker

       }

    }
    return nil

}

func (tgt *TargetDir) ChasingDir()(err error){

    dir, err := os.Open(tgt.Path)

    if err != nil {
        return  err
    }

    dir_content , err := dir.Readdirnames(-1)
    dir.Close()

    for {

        tgt.Marker=actuator.Get_mtime(tgt.Path)

        if (tgt.Marker!=tgt.OldMarker){

           for chan_id :=range tgt.InfoIn {

               tgt.InfoIn[chan_id] <- true

           }

           var current_targets []string

           for chan_id :=range tgt.InfoOut {

               current_targets=append(current_targets,<-tgt.InfoOut[chan_id])

           }

           for cur_id :=range current_targets {

              var found bool

              for prev_id :=range dir_content  {
                  if (dir_content[prev_id]==current_targets[cur_id]) { found=true }

              }
              if (found == false) {
                  new_item_path:=current_targets[cur_id]
                  var new_items = []string {new_item_path}
                  Start(new_items,tgt.MessageChannel)

              }

           }
           dir_content=current_targets
        }

        tgt.OldMarker=tgt.Marker

    }

}


func (tgt *Target) Reporting (){

    tgt.MessageChannel <- tgt.Path

}

func Listen() (messages chan string){


    var test_dir= []string {"/etc/apt"}
    Start(test_dir,messages)
    return





}

func main() {

messages:=Listen()

for {

select{
    case message:=<-messages:fmt.Println(message)
    default:
        time.Sleep(100 * time.Millisecond)
        fmt.Println("No messages")

}

}

}
