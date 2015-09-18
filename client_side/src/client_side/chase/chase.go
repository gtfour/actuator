//package chase
package main

import "client_side/actuator"
import "os"
import "fmt"
import "time"


type Target struct {

    Path string
    Dir string
    OldMarker string
    Marker string
    Modified bool
    EventGroup string
    EventType string
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

    message_channel<-"Starting"

    subdirs:=make(map[string]*TargetDir)

    for id :=range targets {

    file_struct,err:= actuator.Get_md5_file(targets[id])

    if err!=nil {

        dir_struct,err:=actuator.Get_md5_dir(targets[id])

        if err!=nil { continue  } // was a return err


        for subname:=range dir_struct.SubDirs  {


            path :=dir_struct.SubDirs[subname]

            if _, ok := subdirs[path]; ok == false {

                tgt_dir:=&TargetDir{}

                tgt_dir.MessageChannel=message_channel

                tgt_dir.Path=path

                subdirs[path]=tgt_dir
                //go tgt_dir.ChasingDir()
            }



        }


            for file_id :=range dir_struct.Files{

                file_struct:=dir_struct.Files[file_id]

                target:=Target{}
                target.Path=file_struct.Path
                target.OldMarker=string(file_struct.Sum)
                target.MessageChannel=message_channel
                target.InfoIn=make(chan bool,1)
                target.InfoOut=make(chan string, 1)

                if subdir, ok := subdirs[file_struct.Dir]; ok {

                    target.Dir=file_struct.Dir
                    subdir.InfoIn=append(subdir.InfoIn,target.InfoIn)
                    subdir.InfoOut=append(subdir.InfoOut,target.InfoOut)


                }//else {
                //    target.InfoIn = request_channel
                //    target.InfoOut = response_channel
                //}
                go target.ChasingFile()


            }
            //for i:=range subdirs {

           //      subdirs[i].ChasingDir()


            //}

    }else {

      target:=Target{}
      target.Path=targets[id]
      target.OldMarker=string(file_struct.Sum)
      //target.InfoIn = request_channel
      //target.InfoOut = response_channel
      target.MessageChannel=message_channel
      go target.ChasingFile()


    }
    }


    for i:=range subdirs {

        go subdirs[i].ChasingDir()

    }


    return nil

}

func Stop()(err error) {

    return nil
}


func (tgt *Target) ChasingFile() (err error){

    for {

        if (tgt.Dir!="") {
            select {
                case ask_path:= <-tgt.InfoIn:

                    //ask_path:= <-tgt.InfoIn
                    if(ask_path==true) { tgt.InfoOut <- tgt.Path } else { return nil }

                default:

                    if file,err:=actuator.Get_md5_file(tgt.Path);err==nil { tgt.Marker=string(file.Sum) } else { return err }

                    if (tgt.Marker!=tgt.OldMarker){ go tgt.Reporting() }

                    tgt.OldMarker=tgt.Marker

        }

       } else {


          //tgt.MessageChannel<-"chasing file without parent: "+tgt.Path
          if file,err:=actuator.Get_md5_file(tgt.Path);err==nil { tgt.Marker=string(file.Sum) } else { return err }
          if (tgt.Marker!=tgt.OldMarker) { go tgt.Reporting() }

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
    var dir_files []string
    for i:=range dir_content {

        path:=dir_content[i]
        is_dir,err:=actuator.IsDir(path)

        if (err==nil){
            if is_dir==false {

                dir_files=append(dir_files,path)

            }



        }
    }

    for {

        //tgt.MessageChannel<-tgt.Marker+"--"+tgt.OldMarker
        tgt.Marker=actuator.Get_mtime(tgt.Path)

        if (tgt.Marker!=tgt.OldMarker) && (tgt.OldMarker!=""){


           tgt.MessageChannel<-tgt.Marker+"--"+tgt.OldMarker

           if (len(tgt.InfoIn)>0) {

               tgt.MessageChannel<-"channel size :"+string(len(tgt.InfoIn))+"nothing"

               for chan_id :=range tgt.InfoIn {

                   tgt.MessageChannel<-"send name request to childs:"+string(chan_id)

                   tgt.InfoIn[chan_id] <- true

               }

           var current_targets []string

           for chan_id :=range tgt.InfoOut {

               select{
                   case path_value:=<-tgt.InfoOut[chan_id]:
                       current_targets=append(current_targets,path_value)
                   default: continue
                   }

           }

           var new_targets []string

           for cur_id :=range current_targets {

              var found bool

              tgt.MessageChannel<-"cur_targ"+current_targets[cur_id]

              for prev_id :=range dir_files {

                  if (dir_files[prev_id]==current_targets[cur_id]) {

                      tgt.MessageChannel<-"existing: "+dir_files[prev_id] + "new:"+current_targets[cur_id]

                      found=true

                      break

                   }

              }

              if (found == false) {

                  new_item_path:=current_targets[cur_id]
                  new_targets=append(new_targets,new_item_path)

                  //var new_items = []string {new_item_path}

                  //Start(new_items,tgt.MessageChannel)

              }

           }

           if (len(new_targets)>0) {

               var new_items = []string {tgt.Path}
               defer Start(new_items,tgt.MessageChannel)

               for chan_id :=range tgt.InfoIn {

                   tgt.InfoIn[chan_id] <- false

               }

               return nil }


           dir_files=current_targets

          }

        }

        tgt.OldMarker=tgt.Marker

    }

}


func (tgt *Target) Reporting (){

    tgt.MessageChannel <- tgt.Path+"file was modified"

}

func Listen() (messages chan string){


    messages=make(chan string,100)
    var test_dir= []string {"/tmp/test"}
    Start(test_dir,messages)
    return





}

func main() {

messages:=Listen()

for {

select{
    case message:=<-messages:
        fmt.Println(message)
        //time.Sleep(10 * time.Millisecond)
    default:
        time.Sleep(1000 * time.Millisecond)
        fmt.Println("No messages")

}

}

}
