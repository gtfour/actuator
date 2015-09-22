//package chase
package main

import "client_side/actuator"
import "os"
import "fmt"
import "time"
//
//pprof debug
import _ "net/http/pprof"
import "net/http"
//
//


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
            if err!=nil { continue } // was a return err
            //add root dir
            if _, ok := subdirs[targets[id]]; ok == false {
                tgt_dir:=&TargetDir{}
                tgt_dir.MessageChannel=message_channel
                tgt_dir.Path=targets[id]
                subdirs[targets[id]]=tgt_dir
            }
            // 
            for subname:=range dir_struct.SubDirs  {
                path :=dir_struct.SubDirs[subname]
                message_channel <- "subdir :" +path
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
                target.InfoOut=make(chan string,1)
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
            for i:=range subdirs {
                 go subdirs[i].ChasingDir()
            }
    }else{
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

       var inform_about_exit bool

        if (tgt.Dir!="") {

            select {

                case ask_path:= <-tgt.InfoIn:


                    if(ask_path==true) { if (inform_about_exit==true) { tgt.InfoOut <- "|exited|"  } else {  tgt.InfoOut <- tgt.Path  }  } else {  tgt.MessageChannel<-"child is killing self"+tgt.Path  ; return nil }

                default:

                    if file,err:=actuator.Get_md5_file(tgt.Path);err==nil { tgt.Marker=string(file.Sum) } else { inform_about_exit=true  }  //; return err }

                    if (tgt.Marker!=tgt.OldMarker){ go  tgt.Reporting() } else {time.Sleep(10 * time.Millisecond)}

                    tgt.OldMarker=tgt.Marker

        }

       } else {
          //tgt.MessageChannel<-"chasing file without parent: "+tgt.Path
          if file,err:=actuator.Get_md5_file(tgt.Path);err==nil { tgt.Marker=string(file.Sum) } else { /*tgt.InfoOut <- "|exited|"  }  ;*/ return err }
          if (tgt.Marker!=tgt.OldMarker) { go tgt.Reporting() } else {time.Sleep(10 * time.Millisecond)}

                    tgt.OldMarker=tgt.Marker

      }
    }
    return nil
}

func (tgt *TargetDir) ChasingDir()(err error){
   //dup
   dir, err := os.Open(tgt.Path)
   if err != nil {
       return  err
   }
   var dir_content_first []string
   dir_content_first , err = dir.Readdirnames(-1)
   dir.Close()
   // dupdup
   var dir_files_first,dir_subdirs_first []string
   for i:=range dir_content_first {
        path:=dir_content_first[i]
        is_dir,err:=actuator.IsDir(tgt.Path+"/"+path)
        if (err==nil){
                   if is_dir==false {
                    dir_files_first=append(dir_files_first,tgt.Path+"/"+path)

        } else { dir_subdirs_first=append(dir_subdirs_first,tgt.Path+"/"+path) }
        }
    }
           //dupdup
           //dup
    //tgt.OldMarker=actuator.Get_mtime(tgt.Path)
    for {
        tgt.Marker=actuator.Get_mtime(tgt.Path)
        if (tgt.Marker!=tgt.OldMarker) {
           //dup 
           dir, err = os.Open(tgt.Path)
           if err != nil {
           return  err
           }
           var dir_content []string
           dir_content , err = dir.Readdirnames(-1)
           dir.Close()
           // dupdup
           var dir_files,dir_subdirs []string
           for i:=range dir_content {
               path:=dir_content[i]
                   is_dir,err:=actuator.IsDir(tgt.Path+"/"+path)
               if (err==nil){
                   if is_dir==false {
                    dir_files=append(dir_files,tgt.Path+"/"+path)

                } else { dir_subdirs=append(dir_subdirs,tgt.Path+"/"+path) }
              }
           }
           //dupdup
           // dup

               for chan_id :=range tgt.InfoIn {
                   tgt.InfoIn[chan_id] <- true
               }
           var current_targets []string
           var NewInfoIn []chan bool
           var NewInfoOut []chan string

           for chan_id :=range tgt.InfoOut {
               //select{
                   /*case*/ path_value:=<-tgt.InfoOut[chan_id]/*:*/
                       if (path_value!="|exited|") { current_targets=append(current_targets,path_value) ; NewInfoIn=append( NewInfoIn,tgt.InfoIn[chan_id]) ; NewInfoOut=append( NewInfoOut,tgt.InfoOut[chan_id]) }
                   //default: continue
                   //}
           }
           tgt.InfoIn=NewInfoIn
           tgt.InfoOut=NewInfoOut
           var new_targets_files []string
           var new_targets_subdirs []string
           // tratata files
           for cur_id :=range dir_files {
              var found bool
              for prev_id :=range dir_files_first {
                  if (dir_files_first[prev_id]==dir_files[cur_id]) {
                      found=true
                      break
                   }
              }
              if (found == false) {
                  new_item_path:=dir_files[cur_id]
                  new_targets_files=append(new_targets_files,new_item_path)
              }
           }
           dir_files_first=dir_files
           for subdir_id :=range dir_subdirs {
              var found bool
              for prevsubdir_id :=range dir_subdirs_first  {
                  if (dir_subdirs[subdir_id]==dir_subdirs_first[prevsubdir_id]) {
                      found=true
                      break
                   }
              }
              if (found == false) {
                  new_item_path:=dir_subdirs[subdir_id]
                  new_targets_subdirs=append(new_targets_subdirs,new_item_path)
                  target_dir:=&TargetDir{}
                  target_dir.MessageChannel=tgt.MessageChannel
                  target_dir.Path=new_item_path
                  go target_dir.ChasingDir()
              }
           }
           dir_subdirs_first=dir_subdirs
           if (len(new_targets_files)>0) {
               var new_items = []string {tgt.Path}
               //go Start(new_items,tgt.MessageChannel)
               for chan_id :=range tgt.InfoIn {
                   tgt.InfoIn[chan_id] <- false
               }
               go Start(new_items,tgt.MessageChannel)
               return nil }

          tgt.OldMarker=tgt.Marker
        } else {time.Sleep(10 * time.Millisecond)}
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
go func() {
	fmt.Println(http.ListenAndServe("127.0.0.1:6060", nil))
}()

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
