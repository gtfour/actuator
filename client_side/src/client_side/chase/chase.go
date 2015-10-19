//package chase
package main

import "client_side/actuator"
//import "os"
import "fmt"
import "time"
import "path/filepath"
//
//pprof debug
import _ "net/http/pprof"
import "net/http"
//
//
// Before adding TaskMNG


type Target struct {

    Path           string
    Dir            string
    OldMarker      string
    Marker         string
    Modified       bool
    EventGroup     string
    EventType      string
    InfoIn         chan bool
    InfoOut        chan string
    MessageChannel chan string

}

type TargetDir struct {

    Path                 string
    Dir                  string
    OldMarker            string
    Marker               string
    InOutChannelsCreated bool
    InfoIn               chan bool
    InfoOut              chan string
    InfoInArray          []chan bool
    InfoOutArray         []chan string
    MessageChannel       chan   string

}


func Start (targets []string, message_channel chan string , subdirs *map[string]*TargetDir )(err error){
    //request_channel:=make(chan bool)
    //response_channel:=make(chan string)

    message_channel <- "Starting" // message just for information

    //subdirs := make(map[string]*TargetDir) // make subdirs map 

    for id :=range targets {

        file_struct:=&actuator.File{} // create File instance

        err := file_struct.Get_md5_file(targets[id]) // calculate File md5 sum 

        if err.Error()=="is_dir" { // if file is directory

            //fmt.Println(err.Error())

            dir_struct := &actuator.Directory{}

            err := dir_struct.Get_md5_dir(targets[id]) // collect information about included files and directories 

            if err !=nil { continue } // was a return err
            //add root dir

            if _, ok := (*subdirs)[targets[id]]; ok == false { // if global subdirs map does'not contain this item  targets[id] , create and add item to subdirs

                tgt_dir                := &TargetDir{}
                tgt_dir.MessageChannel =  message_channel // bind to main info-channel
                tgt_dir.Path           =  targets[id]
                (*subdirs)[targets[id]]   =  tgt_dir

            }
            // 
            for subname:=range dir_struct.SubDirs  { // iteration of each included subdir

                path := dir_struct.SubDirs[subname]

                if _, ok := (*subdirs)[path]; ok == false { // check global subdir map again and add each included subdir if it is not included yet 

                    tgt_dir                := &TargetDir{}
                    tgt_dir.MessageChannel =  message_channel
                    tgt_dir.Path           =  path
                    (*subdirs)[path]       =  tgt_dir

                //go tgt_dir.ChasingDir()  // i commented this line because it caused chasing directory before assisgning all linked channels to channel Array's
                }
            }
            for file_id :=range dir_struct.Files {

                file_struct            :=  dir_struct.Files[file_id]

                target                 :=  Target{}
                target.Path            =   file_struct.Path
                target.OldMarker       =   string(file_struct.Sum)
                target.MessageChannel  =   message_channel
                target.InfoIn          =   make(chan bool,1)
                target.InfoOut         =   make(chan string,1)

                if subdir, ok := (*subdirs)[file_struct.Dir]; ok { // check Dir field of File struct and try to bind File channel with TargetDir channel Array 

                    target.Dir          =  file_struct.Dir
                    subdir.InfoInArray  =  append(subdir.InfoInArray,target.InfoIn)
                    subdir.InfoOutArray =  append(subdir.InfoOutArray,target.InfoOut)

                }//else {
                //    target.InfoIn = request_channel
                //    target.InfoOut = response_channel
                //}
                go target.ChasingFile()

            }
            //for i:=range subdirs {

           //      message_channel<-"chasing subdir : " + subdirs[i].Path
                 // directory.Dir = filepath.Dir(path)

           //      go subdirs[i].ChasingDir()

           // }
    } else if err == nil {

      // 
      //if err.Error()!="isnt_reg" {
          target                 :=  Target{}
          target.Path            =   targets[id]
          target.OldMarker       =   string(file_struct.Sum)
          //target.InfoIn = request_channel
          //target.InfoOut = response_channel
          target.MessageChannel  =   message_channel
          go target.ChasingFile()
      //}

        }
    }
    for i:=range (*subdirs) {

        //message_channel <- "subdir : " +subdirs[i].Path
        // ебучее говно : поиск родительской директории для этой субдиректории в массиве субдиректорий 

        dir := filepath.Dir(i)

        // ааа бляять - мои мозги !!!! 
        // need to replace dir to i 


        //fmt.Printf(">> BEFORE START : current_dir : %s parent_dir: %s\n",i,dir)

        if parent_dir, ok := (*subdirs)[dir]; ok {

            if (!(*subdirs)[i].InOutChannelsCreated ) {

                (*subdirs)[i].InfoIn          =   make(chan bool,1)
                (*subdirs)[i].InfoOut         =   make(chan string,1)

            }

            (*subdirs)[i].Dir             =   dir
            //fmt.Printf(">> BEFORE START : current_dir : %s parent_dir: %s\n",i,dir)

            parent_dir.InfoInArray          =  append(parent_dir.InfoInArray, (*subdirs)[i].InfoIn)
            parent_dir.InfoOutArray         =  append(parent_dir.InfoOutArray, (*subdirs)[i].InfoOut)

        }
        //go subdirs[i].ChasingDir()

    }

    for i:=range (*subdirs) {

        go (*subdirs)[i].ChasingDir()

    }

    return
}

func Stop()(err error) {
    return nil
}


func (tgt *Target) ChasingFile() (err error){

    for {

       var inform_about_exit bool

        if ( tgt.Dir!="" ) {

            select {

                case ask_path:= <-tgt.InfoIn:

                    if ( ask_path==true ) {

                        if ( inform_about_exit==true ) {

                            tgt.InfoOut  <- "|exited|"

                             _           =  <-tgt.InfoIn /* second signal should be false */

                            return nil  } else {

                                tgt.InfoOut <- tgt.Path  }

                    } else {

                        return nil }

                default:

                    file  :=  &actuator.File{}

                    if err:=file.Get_md5_file(tgt.Path) ; err==nil {

                        tgt.Marker=string(file.Sum) } else {

                        tgt.MessageChannel<-"child is faced with ERROR :" + tgt.Path + "::>>" + err.Error()

                        inform_about_exit=true  }

                    if ( tgt.Marker!=tgt.OldMarker ) {

                        go  tgt.Reporting()

                        tgt.OldMarker=tgt.Marker } else {

                         time.Sleep(200 * time.Millisecond) }

                    //tgt.OldMarker=tgt.Marker

        }

       } else {

          file:=&actuator.File{}

          if err:=file.Get_md5_file(tgt.Path); err == nil {

              tgt.Marker=string(file.Sum) } else {

              /*tgt.InfoOut <- "|exited|"  }  ;*/ return err }

          if (tgt.Marker!=tgt.OldMarker) {

              go tgt.Reporting() ; tgt.OldMarker=tgt.Marker  } else {
              time.Sleep(200 * time.Millisecond)  }
                    //tgt.OldMarker=tgt.Marker
      }
    }
    return nil
}

func (tgt *TargetDir) ChasingDir () (err error){

    var inform_about_exit bool

    tgt.OldMarker , err =  actuator.Get_mtime(tgt.Path)

    if err != nil { return err }

    for {

        tgt.Marker, err  =  actuator.Get_mtime(tgt.Path)

        if err != nil { return err }

        if tgt.Dir != "" {
            select {
                case ask_path:= <-tgt.InfoIn:
                    if ( ask_path==true ) {
                        if ( inform_about_exit == true ) {
                            tgt.InfoOut        <- "|exited|"
                            _                  =  <-tgt.InfoIn /* second signal should be false */
                            return nil

                    } else { tgt.InfoOut <- tgt.Path
                             _           =  <-tgt.InfoIn
                             return nil }
                   } else { return nil }

                default:
            }

        if ( inform_about_exit == true ) {

            var new_items                =  []string { tgt.Path }
            subdirs                      := make(map[string]*TargetDir)
            tgt_new                      := &TargetDir{}
            tgt_new.MessageChannel       =  tgt.MessageChannel
            tgt_new.Path                 =  tgt.Path
            tgt_new.InfoIn               =  tgt.InfoIn
            tgt_new.InfoOut              =  tgt.InfoOut
            tgt_new.Dir                  =  tgt.Dir
            tgt_new.InOutChannelsCreated =  true
            subdirs[tgt.Path]            =  tgt_new

            go Start( new_items, tgt.MessageChannel, &subdirs )

            return nil

        }

        } else if inform_about_exit == true {


            var new_items = []string { tgt.Path }

            subdirs                      := make(map[string]*TargetDir)
            tgt_new                      := &TargetDir{}
            tgt_new.MessageChannel       =  tgt.MessageChannel
            tgt_new.Path                 =  tgt.Path
            tgt_new.InfoIn               =  tgt.InfoIn
            tgt_new.InfoOut              =  tgt.InfoOut
            tgt_new.InOutChannelsCreated =  true
            subdirs[tgt.Path]            =  tgt_new


            go Start( new_items, tgt.MessageChannel, &subdirs )
            return nil

        }

        if ( tgt.Marker != tgt.OldMarker ) {

           for chan_id :=range tgt.InfoInArray {
               tgt.InfoInArray[chan_id] <- true
           }

           var current_targets []string
           var NewInfoInArray  []chan   bool
           var NewInfoOutArray []chan   string

           for chan_id  :=  range tgt.InfoOutArray {

               path_value  :=<-tgt.InfoOutArray[chan_id]

                       if ( path_value      !=  "|exited|" ) {
                           current_targets  =  append( current_targets, path_value )
                           NewInfoInArray   =  append( NewInfoInArray, tgt.InfoInArray[chan_id] )
                           NewInfoOutArray  =  append( NewInfoOutArray, tgt.InfoOutArray[chan_id] )
                       } else {
                           tgt.InfoInArray[chan_id] <- false
                       }
           }

           tgt.InfoInArray  =  NewInfoInArray
           tgt.InfoOutArray =  NewInfoOutArray

           for chan_id := range tgt.InfoInArray {

                   tgt.InfoInArray[chan_id] <- false

           }

           inform_about_exit = true

        } else { time.Sleep( 200 * time.Millisecond )  }
    }
}


func (tgt *Target) Reporting () {

    tgt.MessageChannel <- tgt.Path+"file was modified"

}

func Listen() (messages chan string) {

    target_dir_path             :=  "/proc/net"

    messages                    =   make(chan string,100)
    var test_dir                =   []string { target_dir_path }
    target                      :=  &TargetDir{}
    target.Path                 =   target_dir_path
    target.InfoIn               =   make(chan bool,1)
    target.InfoOut              =   make(chan string,1)
    target.InOutChannelsCreated =   true
    target.MessageChannel       =   messages
    subdirs                     :=  make(map[string]*TargetDir)
    subdirs[target.Path]        =   target


    Start( test_dir, messages, &subdirs )

    return

}

func main() {

    messages:=Listen()

    go func() {
	fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
    }()

    for {

        select{
            case message:=<-messages:
                fmt.Println(message)

            default:
                time.Sleep(1000 * time.Millisecond)
                fmt.Println("No messages")

        }

    }

}
