package chase
//package main

import "client/actuator"
import "client/evebridge"
//import "os"
import "fmt" // for  debug
//import "time"
import "path/filepath"
//import "reflect"
//
//pprof debug
//import _ "net/http/pprof"
//import "net/http"
//
//

type Target struct {

    Path                          string
    Dir                           string
    Prop                          *actuator.Prop // try to use Prop comparing instead of using markers
    InfoIn                        chan bool
    InfoOut                       chan string
    MessageChannel                chan evebridge.CompNotes
    WorkerPool                    *WorkerPool
    InformAboutExit               bool
    KeepChaseWhenDoesNotExist     bool // Do not remove target  from Worker targets array when some error has been caused

}

func ( tgt *Target ) GetDir()  string { return tgt.Dir }
func ( tgt *Target ) GetPath() string { return tgt.Path }
//func ( tgt *Target ) GetProp() *actuator.Prop { return tgt.Prop }
func ( tgt *Target ) GetMessageChannel() chan evebridge.CompNotes {return tgt.MessageChannel}
func ( tgt *Target ) IsReady() bool {return tgt.Prop.Ready }
func ( tgt *Target ) SetReady(state bool)() {tgt.Prop.Ready = state }

type TargetDir struct {

    Target
    InOutChannelsCreated bool
    InfoInArray          []chan bool
    InfoOutArray         []chan string

}

func ( tgt *TargetDir ) GetDir()  string { return tgt.Dir }
func ( tgt *TargetDir ) GetPath() string { return tgt.Path }
//func ( tgt *TargetDir ) GetProp() *actuator.Prop { return tgt.Prop }
func ( tgt *TargetDir ) GetMessageChannel() chan evebridge.CompNotes {return tgt.MessageChannel}
func ( tgt *TargetDir ) IsReady() bool {return tgt.Prop.Ready }
func ( tgt *TargetDir ) SetReady(state bool)() {tgt.Prop.Ready = state }


func Start (targets []string, message_channel chan evebridge.CompNotes ,wp *WorkerPool, subdirs *map[string]*TargetDir )(err error){

    // Bug was found : when i passing an file name not a dir name to func to func Start 
    // nil point dereference error is causing . It happens because subdirs is nil ( i suppose )
    // Solution: Seems that problem was caused when i was trying to get text message of error which was nil . "is_dir"

    //message_channel <- "Starting"

    for id :=range targets {

        fstruct       := &actuator.File{} // create File instance
        fstruct.Prop  =  actuator.GetProp(targets[id],SAFE_OPENING_MODE) // calculate File md5 sum 

        if fstruct.Prop.IsDir == true  { // if file is directory
            dir_struct := &actuator.Directory{}
            err := dir_struct.GetHashSumDir(targets[id],SAFE_OPENING_MODE) // collect information about included files and directories 
            if err !=nil { continue } // was a return err
            if _, ok := (*subdirs)[targets[id]]; ok == false { // if global subdirs map does'not contain this item  targets[id] , create and add item to subdirs
                tgt_dir                   := TargetDir{}
                tgt_dir.MessageChannel    =  message_channel // bind to main info-channel
                tgt_dir.Path              =  targets[id]
                tgt_dir.WorkerPool        =   wp // Deirz@golang.cjr advice  
                (*subdirs)[targets[id]]   =  &tgt_dir
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
                //fmt.Printf("\ntgt path :%s\n",file_struct.Path)
                prop               :=  actuator.GetProp(file_struct.Path, SAFE_OPENING_MODE)

                if prop.Error == true { continue }

                target                 :=  Target{} //  I have to find difference between Target{} and &Target{}
                target.Prop            =  prop

                target.Path            =   file_struct.Path
                target.MessageChannel  =   message_channel
                target.WorkerPool      =   wp
                target.InfoIn          =   make(chan bool,1)
                target.InfoOut         =   make(chan string,1)

                if subdir, ok := (*subdirs)[file_struct.Dir]; ok { // check Dir field of File struct and try to bind File channel with TargetDir channel Array 

                    target.Dir          =  file_struct.Dir
                    subdir.InfoInArray  =  append(subdir.InfoInArray,target.InfoIn)
                    subdir.InfoOutArray =  append(subdir.InfoOutArray,target.InfoOut)

                }
                if (wp==nil) {/*fmt.Printf("wp is nill dir files ")*/}
                wp.AppendTarget(&target)

            }

    } else if fstruct.Prop.IsRegular == true  {

          target                 :=  Target{}
          target.Path            =   targets[id]
          prop               :=  actuator.GetProp(targets[id], SAFE_OPENING_MODE)
          if prop.Error == true  { continue }
          target.Prop            =   prop
          target.WorkerPool      =   wp // new 02-11-2015 03:00

	  target.MessageChannel  =   message_channel
	  if (wp==nil) {/*fmt.Printf("wp is nill single files ")*/}
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
        prop      :=  actuator.GetProp(target_subdir.Path, SAFE_OPENING_MODE)
        if prop.Error == true { continue }

        target_subdir.Prop  = prop
        //go (*subdirs)[i].Chasing()
        if (wp==nil) {/*fmt.Printf("wp is nill dir subdirs ")*/}
        wp.AppendTarget(target_subdir)

    }

    return
}

func Stop()(err error) {
    return nil
}


func (tgt *Target) Chasing(mode int) (err error){

    fmt.Printf("\nChasing:file")

    //for {

        if ( tgt.Dir!="" ) {

            select {

                case <-tgt.InfoIn:

                            return nil


                default:

                    actual_prop  :=   actuator.GetProp(tgt.Path,mode)

                    if actual_prop.Error == true {
                        error_field:=evebridge.CompNote{Field:"Error",Before:"false",After:"true"}
                        cnote      :=evebridge.CompNotes{Path:tgt.Path}
                        cnote.List = append(cnote.List, error_field)
                        tgt.MessageChannel <- cnote
                        tgt.InformAboutExit=true
                        return err

                    }


                    //if ( reflect.DeepEqual(actual_prop, tgt.Prop) == false ) {
                    if comparison_notes:=actuator.CompareProp(actual_prop, tgt.Prop, tgt.Path ); len(comparison_notes.List)>0 {

                        //go  tgt.Reporting()
                        tgt.MessageChannel <- comparison_notes

                        tgt.Prop = actual_prop }


                    }

       } else {
           actual_prop  :=   actuator.GetProp(tgt.Path,mode)

           if actual_prop.Error == true {
               error_field:=evebridge.CompNote{Field:"Error",Before:"false",After:"true"}
               cnote      :=evebridge.CompNotes{Path:tgt.Path}
               cnote.List = append(cnote.List, error_field)
               tgt.MessageChannel <- cnote
               tgt.InformAboutExit=true
               return err

           }


           //if ( reflect.DeepEqual( actual_prop, tgt.Prop ) == false ) {
           if comparison_notes:=actuator.CompareProp(actual_prop, tgt.Prop, tgt.Path ); len(comparison_notes.List)>0 {

               //go  tgt.Reporting()
               tgt.MessageChannel <- comparison_notes

               tgt.Prop = actual_prop

           }


      }
    return nil
}

func (tgt *TargetDir) Chasing (mode int) (err error){

        fmt.Printf("Chasing:dir")


        actual_prop  :=  actuator.GetProp(tgt.Path, mode)

        if actual_prop.Error == true { /*fmt.Printf("\nError during opening %s\n",tgt.Path)*/ }

        if (tgt.WorkerPool==nil) {/*fmt.Printf("%s wp is nil",tgt.Path)*/}

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
                tgt.Prop=actual_prop

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
            tgt.Prop=actual_prop

            go Start( new_items, tgt.MessageChannel, tgt.WorkerPool, &subdirs )

            return nil

        }

        //if ( reflect.DeepEqual( actual_prop, tgt.Prop ) == false ) {
        if comparison_notes:=actuator.CompareProp(actual_prop, tgt.Prop, tgt.Path ); len(comparison_notes.List)>0 {
           tgt.MessageChannel <- comparison_notes
           tgt.Prop = actual_prop

           for chan_id :=range tgt.InfoInArray {
               tgt.InfoInArray[chan_id] <- true
           }

           tgt.InformAboutExit = true

        }
    return nil
}


//func (tgt *Target) Reporting () {

//    tgt.MessageChannel <- tgt.Path+"file was modified"

//}

func Listen( path string,  messages chan evebridge.CompNotes , wp WorkerPool )(err error) {

    target_dir_path             :=  path

    // // messages                    =   make(chan string,100)
    // // wp                          :=  WPCreate()
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

    return nil

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
