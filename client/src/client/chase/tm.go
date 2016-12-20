package chase


//import "fmt"
import "time"
import "math/rand"
import "jumper/actuator"
import "client/majesta"
// Have to implement carousel method


var TGT_PER_GR int64                       = 50 // if FILES_PER_GR is very big - TargetsCount type should be modified 
var TIMEOUT_MS              time.Duration  = 800
var INHIBITION_TIMEOUT      time.Duration  = 1000
var LOG_CHANNEL_TIMEOUT_MS  time.Duration  = 1000


var EMPTY_OPENING_MODE int = 0
var LAZY_OPENING_MODE  int = 01
var SAFE_OPENING_MODE  int = 02


type AbstractTarget interface {
    GetDir()            string
    Chasing(int)        error
    GetPath()           string
    GetMessageChannel() chan majesta.CompNotes
    SetReady            (bool)()
    IsReady()           bool
    GetSelfProp         ()(*actuator.Prop)
    CloseFd()()
    SetOpeningMode(int)()
    GetOpeningMode()(int)
    AskInitialCheck()()
    GetRemove()(bool)
    ToRemove()()
}

type WorkerPool struct {
    Workers           []*Worker
    WKillers          []chan bool
    ReadyTargets      chan AbstractTarget
    RunningTargets    chan AbstractTarget
    SuspendedTargets  chan AbstractTarget
    ManageChannel     chan *ManMessage
    Targets           []AbstractTarget
}


type Worker struct {
    WorkerPool    *WorkerPool
    Targets       []AbstractTarget
    Id            int32
    TargetsCount  int32
    Stop          chan bool
}

type ManMessage struct {

    operation string
    tgt       AbstractTarget
    path      string


}

func ( w *Worker ) Start ()  {
        for {
            select {
            case <-w.Stop:
                break
            case tgt :=<-w.WorkerPool.ReadyTargets:
                go func() {
                        opening_mode := tgt.GetOpeningMode()
                        if opening_mode == LAZY_OPENING_MODE || opening_mode == EMPTY_OPENING_MODE  {
                            w.WorkerPool.RunningTargets <- tgt
                            _                         = tgt.Chasing(LAZY_OPENING_MODE) //should be  light file opening
                            tgt.SetReady(true)
                        } else {
                            _                         = tgt.Chasing(SAFE_OPENING_MODE)
                            tgt.SetReady(true)
                            w.WorkerPool.ReadyTargets <- tgt // has been uncommented
                        }
                }()
            }
        }
}

func ( w *Worker ) Append ( tgt AbstractTarget ) {

        w.Targets = append(w.Targets,tgt)
}

func WPCreate () (wp WorkerPool) {

    wp.Workers          = make([]*Worker, 0)
    wp.ReadyTargets     = make(chan AbstractTarget,100)
    wp.RunningTargets   = make(chan AbstractTarget,100)
    wp.SuspendedTargets = make(chan AbstractTarget,100)
    wp.ManageChannel    = make(chan *ManMessage,   300)
    wp.AddWorker()
    wp.AddWorker()
    go wp.Juggle()
    go wp.Management()

    return

}




func ( wp *WorkerPool ) Stop () {
    killers := wp.WKillers
    for i:= range killers {
        killers[i] <- true
    }
}




func ( wp *WorkerPool ) Juggle () {
        go func(){
            for {
                select {
                    case tgt :=<-wp.SuspendedTargets:
                        if tgt.IsReady() == true {
                            tgt.SetOpeningMode(LAZY_OPENING_MODE)
                            tgt.SetReady(false)
                            wp.ReadyTargets <- tgt
                        } else {
                            if tgt.GetOpeningMode() == LAZY_OPENING_MODE {
                                tgt.CloseFd()
                                tgt.SetOpeningMode(SAFE_OPENING_MODE)
                                tgt.SetReady(false)
                                wp.ReadyTargets <- tgt
                            }

                        }
                    }
                }
        }()
        for {
            select {
                case tgt := <-wp.RunningTargets:
                        if tgt.GetRemove() == false {
                            if tgt.IsReady() == true {
                                wp.ReadyTargets     <- tgt
                            } else {
                                wp.SuspendedTargets <- tgt
                            }
                        }
                default:
                    time.Sleep( INHIBITION_TIMEOUT * time.Millisecond )
            }
        }
}



func (wp *WorkerPool)  AddWorker()(){

    w           :=   &Worker{ WorkerPool:wp }
    rand.Seed( time.Now().UTC().UnixNano())
    w.Id        =   rand.Int31()
    w.Stop      =   make(chan bool)
    wp.Workers  =   append(wp.Workers, w)
    go w.Start()

}

func ( wp *WorkerPool ) AppendTarget ( tgt AbstractTarget ) () {

    wp.ManageChannel<-&ManMessage{operation:"add",tgt:tgt}
}

func ( wp *WorkerPool ) RemoveTarget ( path string ) () {

    wp.ManageChannel<-&ManMessage{operation:"remove",path:path}

}

func ( wp *WorkerPool ) Management () {

    for {
        select {

            case task := <-wp.ManageChannel:
                switch task.operation {
                    case "add":
                        var tgt_exists bool
                        for existing_tgt := range wp.Targets {
                            existing_tgt    := wp.Targets[existing_tgt]
                            if task.tgt.GetPath() == existing_tgt.GetPath()  { tgt_exists=true ; break }
                        }
                        if tgt_exists == false {
                            wp.Targets = append(wp.Targets, task.tgt)
                            task.tgt.AskInitialCheck()
                            wp.ReadyTargets <- task.tgt
                        }
                    case "remove":
                        var tgt_id int
                        for i := range wp.Targets {
                            existing_target:= wp.Targets[i]
                            if  (existing_target.GetPath()==task.path) {
                                existing_target.ToRemove()
                                tgt_id = i
                                break
                            }
                        }
                        wp.Targets = append(wp.Targets[:tgt_id], wp.Targets[tgt_id+1:]...)
                }
        }
    }
}
