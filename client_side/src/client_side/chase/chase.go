package chase
//package main

import "client_side/actuator"
//import "os"
import "fmt" // for  debug
//import "time"
import "path/filepath"
//
//pprof debug
//import _ "net/http/pprof"
//import "net/http"
//
//

type Target struct {
    Path                          string
    Dir                           string
    OldMarker                     string
    Marker                        string
    InfoIn                        chan bool
    InfoOut                       chan string
    MessageChannel                chan string
    WorkerPool                    *WorkerPool
    InformAboutExit               bool
    KeepChaseWhenDoesNotExist     bool
    MarkerGetttingModeIsMtime     bool // "mtime" or "md5sum"

}

func ( tgt *Target ) GetDir()  string { return tgt.Dir }
func ( tgt *Target ) GetPath() string { return tgt.Path }

type TargetDir struct {

    Target
    InOutChannelsCreated bool
    InfoInArray          []chan bool
    InfoOutArray         []chan string

}

func ( tgt *TargetDir ) GetDir()  string { return tgt.Dir }
func ( tgt *TargetDir ) GetPath() string { return tgt.Path }


func Start (targets []string, message_channel chan string ,wp *WorkerPool, subdirs *map[string]*TargetDir )(err error){

    // Bug was found : when i passing an file name not a dir name to func to func Start 
    // nil point dereference error is causing . It happens because subdirs is nil ( i suppose )
    // Solution: Seems that problem was caused when i was trying to get text message of error which was nil . "is_dir"

    message_channel <- "Starting"

    for id :=range targets {

        file_struct:=&actuator.File{} // create File instance
        err := file_struct.Get_md5_file(targets[id]) // calculate File md5 sum 
        if err==actuator.Is_dir_error { // if file is directory
            dir_struct := &actuator.Directory{}
            err := dir_struct.Get_md5_dir(targets[id]) // collect information about included files and directories 
            if err !=nil { continue } // was a return err
            if _, ok := (*subdirs)[targets[id]]; ok == false { // if global subdirs map does'not contain this item  targets[id] , create and add item to subdirs
                tgt_dir                   := &TargetDir{}
                tgt_dir.MessageChannel    =  message_channel // bind to main info-channel
                tgt_dir.Path              =  targets[id]
                tgt_dir.WorkerPool        =   wp // Deirz@golang.cjr advice  
                (*subdirs)[targets[id]]   =  tgt_dir
            }
            // 
            for subname:=range dir_struct.SubDirs  { // iteration of each included subdir
                path := dir_struct.SubDirs[subname]
                if _, ok := (*subdirs)[path]; ok == false { // check global subdir map again and add each included subdir if it is not included yet 
                    tgt_dir                 :=  TargetDir{} // try to change &TargetDir{} to TargetDir{}
                    tgt_dir.MessageChannel  =   message_channel
                    tgt_dir.WorkerPool      =   wp
                    tgt_dir.Path            =   path
                    (*subdirs)[path]        =   &tgt_dir
                }
            }
            for file_id :=range dir_struct.Files {

                file_struct            :=  dir_struct.Files[file_id]

                target                 :=  Target{} //  I have to find difference between Target{} and &Target{}

                if file_struct.MarkerGetttingModeIsMtime == true {
                    target.MarkerGetttingModeIsMtime = true
                }

                target.Path            =   file_struct.Path
                target.OldMarker       =   string(file_struct.Sum)
                target.MessageChannel  =   message_channel
                target.WorkerPool      =   wp
                target.InfoIn          =   make(chan bool,1)
                target.InfoOut         =   make(chan string,1)

                if subdir, ok := (*subdirs)[file_struct.Dir]; ok { // check Dir field of File struct and try to bind File channel with TargetDir channel Array 

                    target.Dir          =  file_struct.Dir
                    subdir.InfoInArray  =  append(subdir.InfoInArray,target.InfoIn)
                    subdir.InfoOutArray =  append(subdir.InfoOutArray,target.InfoOut)

                }
                if (wp==nil) {fmt.Printf("wp is nill dir files ")}
                wp.AppendTarget(&target)

            }

    } else if err == nil || err==actuator.Have_to_switch_to_mtime  {

          target                 :=  Target{}
          target.Path            =   targets[id]
          if file_struct.MarkerGetttingModeIsMtime == true {
              target.MarkerGetttingModeIsMtime = true
          }
          target.OldMarker       =   string(file_struct.Sum)
          target.WorkerPool      =   wp // new 02-11-2015 03:00

	  target.MessageChannel  =   message_channel
	  if (wp==nil) {fmt.Printf("wp is nill single files ")}
          wp.AppendTarget(&target)

        }
    }
    for i:=range (*subdirs) {

        dir := filepath.Dir(i)

        if parent_dir, ok := (*subdirs)[dir]; ok {

            if (!(*subdirs)[i].InOutChannelsCreated ) {

                (*subdirs)[i].InfoIn          =   make(chan bool,1)
                (*subdirs)[i].InfoOut         =   make(chan string,1)

            }

            (*subdirs)[i].Dir             =   dir

            parent_dir.InfoInArray          =  append(parent_dir.InfoInArray, (*subdirs)[i].InfoIn)
            parent_dir.InfoOutArray         =  append(parent_dir.InfoOutArray, (*subdirs)[i].InfoOut)

        }

    }

    for i:=range (*subdirs) {

        target_subdir := (*subdirs)[i]

        target_subdir.OldMarker , err =  actuator.Get_mtime(target_subdir.Path) // this code has been moved here from top of Chasing()
        if err != nil { continue }
        //go (*subdirs)[i].Chasing()
        if (wp==nil) {fmt.Printf("wp is nill dir subdirs ")}
        wp.AppendTarget(target_subdir)

    }

    return
}

func Stop()(err error) {
    return nil
}


func (tgt *Target) Chasing() (err error){

    //for {

        if ( tgt.Dir!="" ) {

            select {

                case <-tgt.InfoIn:

                            return nil


                default:

                    file  :=  &actuator.File{}
                    var marker string
                    if tgt.MarkerGetttingModeIsMtime == true {
                        marker,err= actuator.Get_mtime(tgt.Path)
                    } else {
                        err=file.Get_md5_file(tgt.Path)
                        marker=string(file.Sum)

                    }
                    if err!= nil && err!=actuator.Have_to_switch_to_mtime {
                        tgt.MessageChannel<-"child gets this ERROR :" + tgt.Path + ":>" + err.Error()+"|"
                        tgt.InformAboutExit=true
                        return err

                    } else if err==actuator.Have_to_switch_to_mtime {
                        tgt.MarkerGetttingModeIsMtime = true
                    }
                    tgt.Marker = marker


                    if ( tgt.Marker!=tgt.OldMarker ) {

                        fmt.Printf("\nMarkers are different %s\n",tgt.Path)
                        go  tgt.Reporting()

                        tgt.OldMarker=tgt.Marker }



        }

       } else {

          file:=&actuator.File{}
          var marker string

          if tgt.MarkerGetttingModeIsMtime == true {
              marker,err= actuator.Get_mtime(tgt.Path)
          } else {
              err=file.Get_md5_file(tgt.Path)
              marker=string(file.Sum)
          }
          if err!= nil && err!=actuator.Have_to_switch_to_mtime {
              tgt.MessageChannel<-"child gets this ERROR :" + tgt.Path + ":>" + err.Error()+"|"
              tgt.InformAboutExit=true
              return err
           } else if err==actuator.Have_to_switch_to_mtime {
               tgt.MarkerGetttingModeIsMtime = true
           }
          tgt.Marker = marker



          if (tgt.Marker!=tgt.OldMarker) {
              fmt.Printf("\nMarkers are different %s\n",tgt.Path)

              go tgt.Reporting() ; tgt.OldMarker=tgt.Marker  }
      }
    return nil
}

func (tgt *TargetDir) Chasing () (err error){


        tgt.Marker, err  =  actuator.Get_mtime(tgt.Path)

        if err!= nil { fmt.Printf("\nError during opening %s\n",tgt.Path) }

        if (tgt.WorkerPool==nil) {fmt.Printf("%s wp is nil",tgt.Path)}

        if err != nil { return err }

        if tgt.Dir != "" {
            select {
                case <-tgt.InfoIn:
                            return nil
                default:
            }

        if ( tgt.InformAboutExit == true ) {

            var new_items                =  []string { tgt.Path }
            subdirs                      := make(map[string]*TargetDir)
            tgt_new                      := &TargetDir{}
            tgt_new.MessageChannel       =  tgt.MessageChannel
            tgt_new.Path                 =  tgt.Path
            tgt_new.InfoIn               =  tgt.InfoIn
            tgt_new.InfoOut              =  tgt.InfoOut
            tgt_new.Dir                  =  tgt.Dir
            tgt_new.InOutChannelsCreated =  true
            tgt_new.WorkerPool           =  tgt.WorkerPool
            subdirs[tgt.Path]            =  tgt_new

            // watafa possible mistake cause  was found . it is tgt.InformAboutExit 
            tgt.InformAboutExit = false
            // second possible mistake cause
            tgt.OldMarker=tgt.Marker

            go Start( new_items, tgt.MessageChannel, tgt.WorkerPool, &subdirs )

            return nil

        }

        } else if tgt.InformAboutExit == true {


            var new_items = []string { tgt.Path }

            subdirs                      := make(map[string]*TargetDir)
            tgt_new                      := &TargetDir{}
            tgt_new.MessageChannel       =  tgt.MessageChannel
            tgt_new.Path                 =  tgt.Path
            tgt_new.InfoIn               =  tgt.InfoIn
            tgt_new.InfoOut              =  tgt.InfoOut
            tgt_new.InOutChannelsCreated =  true
            tgt_new.WorkerPool           =  tgt.WorkerPool
            subdirs[tgt.Path]            =  tgt_new

            // watafa possible mistake was found . it is tgt.InformAboutExit 
            tgt.InformAboutExit = false
            // second possible mistake cause
            tgt.OldMarker=tgt.Marker

            go Start( new_items, tgt.MessageChannel, tgt.WorkerPool, &subdirs )

            return nil

        }

        if ( tgt.Marker!=tgt.OldMarker ) {
           fmt.Printf("\nMarkers are different %s\n",tgt.Path)

           for chan_id :=range tgt.InfoInArray {
               tgt.InfoInArray[chan_id] <- true
           }

           tgt.InformAboutExit = true

        }
    return nil
}


func (tgt *Target) Reporting () {

    tgt.MessageChannel <- tgt.Path+"file was modified"

}

func Listen( path string ) ( messages chan string ) {

    target_dir_path             :=  path

    messages                    =   make(chan string,100)
    wp                          :=  WPCreate()
    var test_dir                =   []string { target_dir_path }
    target                      :=  &TargetDir{}
    target.Path                 =   target_dir_path
    //target.SelfType             =   "dir"
    target.InfoIn               =   make(chan bool,1)
    target.InfoOut              =   make(chan string,1)
    target.InOutChannelsCreated =   true
    target.MessageChannel       =   messages
    target.WorkerPool           =   &wp // Deirz@golang.cjr advice
    subdirs                     :=  make(map[string]*TargetDir)
    subdirs[target.Path]        =   target



    Start( test_dir, messages, &wp, &subdirs )

    return

}

/*func main() {

    messages:=Listen()

    go func() {
	fmt.Println(http.ListenAndServe("0.0.0.0:6060", nil))
    }()

    for {

        select{
            case message:=<-messages:
                fmt.Println(message)

            default:
                time.Sleep( LOG_CHANNEL_TIMEOUT_MS * time.Millisecond)
                fmt.Println("No messages")

        }

    }

}*/
