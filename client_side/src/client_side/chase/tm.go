package chase


import "fmt"
import "time"
import "math/rand"


var TGT_PER_GR int64                       = 50 // if FILES_PER_GR is very big - TargetsCount type should be modified 
var TIMEOUT_MS              time.Duration  = 200
var LOG_CHANNEL_TIMEOUT_MS  time.Duration  = 1000


type AbstractTarget interface {
    GetDir()  string
    Chasing() error
    GetPath() string
}

type WorkerPool struct {
    Workers         []*Worker
    WKillers        []chan bool
}


type Worker struct {
    Targets       []AbstractTarget
    Id            int32
    TargetsCount  int32
    Stop          chan bool
    //TargetDirs  []*TargetDir
}

func ( w *Worker ) Start ()  {
    for {
        //fmt.Printf("\n<-- Worker is working %d-->\n",w.Id)
        select {

            case <-w.Stop:

                return

            default:

                for tgt := range w.Targets {

                    if w.Targets[tgt].Chasing!=nil {

                        w.Targets[tgt].Chasing()

                    }

                }
        }
        time.Sleep( TIMEOUT_MS * time.Millisecond )
    }
}

func ( w *Worker ) Append ( tgt AbstractTarget ) {

        w.Targets = append(w.Targets,tgt)
}

func WPCreate () (wp WorkerPool) {

    wp.Workers  = make([]*Worker, 0)
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


func (wp *WorkerPool)  AddWorker()(){



    fmt.Printf("\n -- Create new worker\n")
    w           :=   &Worker{}
    rand.Seed( time.Now().UTC().UnixNano())
    w.Id        =   rand.Int31()
    w.Stop      =   make(chan bool)
    wp.Workers  =   append(wp.Workers, w)
    go w.Start()


}

func ( wp *WorkerPool ) AppendTarget ( tgt AbstractTarget ) () {

    fmt.Printf("\n Appending target %s \n",tgt.GetPath())

    var tgt_replaced  bool
    var all_tgt_count int64

    for w:= range wp.Workers {

        worker := wp.Workers[w]

            for wtgt := range worker.Targets {

                worker_target_dir := worker.Targets[wtgt]
                //fmt.Printf("\n tgt.Dir %s     worker_target_dir.Path %s\n",tgt.GetDir(),worker_target_dir.GetPath())
                if tgt.GetPath() == worker_target_dir.GetPath() { worker_target_dir=tgt  ; tgt_replaced=true  ; break }
                all_tgt_count=all_tgt_count+1

            }


    }
    if tgt_replaced == false {
        var rand_digit int32
        rand.Seed( time.Now().UTC().UnixNano())
        rand_digit = rand.Int31n(int32(len(wp.Workers)))
        wp.Workers[rand_digit].Append(tgt)
    }

    average_tgt_per_worker:=all_tgt_count/int64(len(wp.Workers))
    if average_tgt_per_worker > TGT_PER_GR { wp.AddWorker() }

}
