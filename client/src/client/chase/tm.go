package chase


import "fmt"
import "time"
import "math/rand"
//import "client/actuator"
import "client/evebridge"
// Have to implement carousel method


var TGT_PER_GR int64                       = 50 // if FILES_PER_GR is very big - TargetsCount type should be modified 
var TIMEOUT_MS              time.Duration  = 800
var LOG_CHANNEL_TIMEOUT_MS  time.Duration  = 1000
var LAZY_OPENING_MODE int = 01
var SAFE_OPENING_MODE int = 02


type AbstractTarget interface {
    GetDir()            string
    Chasing(string)     error
    GetPath()           string
    //GetProp()           *actuator.Prop
    GetMessageChannel() chan evebridge.CompNotes
    SetReady            (bool)()
    IsReady()           bool
    //CloseFd()           error
}

type WorkerPool struct {
    Workers         []*Worker
    WKillers        []chan bool
    ReadyTargets    chan AbstractTarget
    RunningTargets  chan AbstractTarget
}


type Worker struct {
    WorkerPool    *WorkerPool
    Targets       []AbstractTarget
    Id            int32
    TargetsCount  int32
    Stop          chan bool
    //TargetDirs  []*TargetDir
}

func ( w *Worker ) Start ()  {
    //ticker := time.NewTicker(TIMEOUT_MS * time.Millisecond)
    //for _ = range ticker.C {
        //fmt.Printf("\n<-- Worker is working %d-->\n",w.Id)
        for {
            select {
            case <-w.Stop:
                //ticker.Stop()
                break
            case tgt :=<-w.WorkerPool.ReadyTargets:
                go func() {
                    if tgt.IsReady() == true{
                        tgt.SetReady(false)
                        w.WorkerPool.RunningTargets <- tgt
                        _                         = tgt.Chasing(LAZY_OPENING_MODE) //should be  light file opening
                        tgt.SetReady(true)
                    } else {
                        w.WorkerPool.RunningTargets <- tgt
                        _                         = tgt.Chasing(SAFE_OPENING_MODE)
                        tgt.SetReady(true)
                    }
                    //w.WorkerPool.ReadyTargets   <- tgt
                }()
            //
            /*default:
                var unused_tgt_numbers []int // array for store tgt numbers whose should be removed from w.Targets
                targets_count:=len(w.Targets)
                for tgt := range w.Targets {
                    if w.Targets[tgt].Chasing != nil {
                        err:=w.Targets[tgt].Chasing("lazy")
                        if err != nil { unused_tgt_numbers=append(unused_tgt_numbers,tgt) }
                    }
                }
                // remove tgt's whose Chasing was returned with err!=nil  w.Targets
                for i := range unused_tgt_numbers {
                    tgt_num:=unused_tgt_numbers[i]
                    if (targets_count == len(w.Targets) ) {
                        w.Targets = append(w.Targets[:tgt_num], w.Targets[tgt_num+1:]...)
                    }
                }
              */
              //
            }
        }
        //time.Sleep( TIMEOUT_MS * time.Millisecond )
    //}
}

func ( w *Worker ) Append ( tgt AbstractTarget ) {

        w.Targets = append(w.Targets,tgt)
}

func WPCreate () (wp WorkerPool) {

    wp.Workers        = make([]*Worker, 0)
    wp.ReadyTargets   = make(chan AbstractTarget,100 )
    wp.RunningTargets = make(chan AbstractTarget,100 )
    go wp.Juggle()
    // try to create two workers instead of one
    wp.AddWorker()
    wp.AddWorker()

    return

}

func ( wp *WorkerPool ) Stop () {
    killers := wp.WKillers
    for i:= range killers {
        killers[i] <- true
    }
}
func ( wp *WorkerPool ) Juggle () {
    ticker := time.NewTicker(TIMEOUT_MS * time.Millisecond)
    //for _ = range ticker.C {
        fmt.Printf("\n--Juggling--\n")
        SuspendedTargets := make(chan AbstractTarget,100)
        go func(){
            ticker := time.NewTicker(TIMEOUT_MS * time.Millisecond)
            for _ = range ticker.C {
                select {
                    case tgt :=<-SuspendedTargets:
                        if tgt.IsReady() == true {
                            wp.ReadyTargets <- tgt
                        } else {

                        }
                    }

             }
        }()
        for {
            select {
                case tgt := <-wp.RunningTargets:
                    if tgt.IsReady() == true {
                            wp.ReadyTargets     <- tgt
                    } else {
                            wp.SuspendedTargets <- tgt
                    }
                default:
                    //fmt.Printf("\nnone")
            }
        }
    //}
}



func (wp *WorkerPool)  AddWorker()(){

    //fmt.Printf("\n -- Create new worker\n")
    w           :=   &Worker{ WorkerPool:wp }
    rand.Seed( time.Now().UTC().UnixNano())
    w.Id        =   rand.Int31()
    w.Stop      =   make(chan bool)
    wp.Workers  =   append(wp.Workers, w)
    go w.Start()

}

func ( wp *WorkerPool ) AppendTarget ( tgt AbstractTarget ) () {

     wp.ReadyTargets <- tgt

    //fmt.Printf("\n Appending target %s \n",tgt.GetPath())
    //bug has been found : can't add targets more than 2 worker * TGT_PER_GR
    // lol, seems that linked to message channel size

    /*
    var tgt_replaced  bool
    var all_tgt_count int64

    for w:= range wp.Workers {

        worker := wp.Workers[w]

            for wtgt := range worker.Targets {

                worker_target := worker.Targets[wtgt]
                //fmt.Printf("\n tgt.Dir %s     worker_target_dir.Path %s\n",tgt.GetDir(),worker_target_dir.GetPath())
                if tgt.GetPath() == worker_target.GetPath() { worker_target=tgt  ; tgt_replaced=true }   //; break }
                all_tgt_count=all_tgt_count+1

            }


    }
    if tgt_replaced == false {
        var rand_digit int32
        rand.Seed( time.Now().UTC().UnixNano())
        rand_digit = rand.Int31n(int32(len(wp.Workers)))
        wp.Workers[rand_digit].Append(tgt)

        //tgt.GetMessageChannel() <- actuator.Initial(tgt.GetProp(), tgt.GetPath())
    }

    average_tgt_per_worker:=all_tgt_count/int64(len(wp.Workers))
    if average_tgt_per_worker > TGT_PER_GR { wp.AddWorker() }
    */

}
