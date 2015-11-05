package chase


import "fmt"
import "time"
import "math/rand"


var FILES_PER_GR                           = 1000 // if FILES_PER_GR is very big - TargetsCount type should be modified 
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
        fmt.Printf("\n<-- Worker is working %d-->\n",w.Id)
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

    //if tgt.GetType() == "dir" {
        var tgt_replaced bool
        for wtgt:= range w.Targets {

            // have mistake . Also need to check all workers in wp pool
            if w.Targets[wtgt].GetPath() == tgt.GetPath() { w.Targets[wtgt] = tgt ; tgt_replaced = true }

        }
        if tgt_replaced == false { w.Targets = append(w.Targets,tgt) }
}

func WPCreate () (wp WorkerPool) {

    wp.Workers  = make([]*Worker, 0)
    wp.AddWorker()

    return

}

func ( wp *WorkerPool ) Stop () {

    killers := wp.WKillers

    for i:= range killers {

        killers[i] <- true

    }

}


func (wp *WorkerPool)  AddWorker()(w *Worker){



    fmt.Printf("\nCreate new worker\n")
    w           =   &Worker{}
    rand.Seed( time.Now().UTC().UnixNano())
    w.Id        =   rand.Int31()
    w.Stop      =   make(chan bool)
    wp.Workers  =   append(wp.Workers, w)
    go w.Start()
    return


}

func ( wp *WorkerPool ) AppendTarget ( tgt AbstractTarget ) () {

    fmt.Printf("\n Appending target %s \n",tgt.GetPath())

    var create_new_worker bool

    for w:= range wp.Workers {

        worker := wp.Workers[w]

        for wtgt := range worker.Targets {

            worker_target_dir := worker.Targets[wtgt]
            fmt.Printf("\n tgt.Dir %s     worker_target_dir.Path %s\n",tgt.GetDir(),worker_target_dir.GetPath())
            if tgt.GetDir() == worker_target_dir.GetPath() { // HAVE to add check tgt.GetPath == worker_target_dir.GetDir()
                create_new_worker = true
                break
            }

        }

       if create_new_worker == false { worker.Append(tgt) ; break  } else  { create_new_worker=false }

       if w == len(wp.Workers) { create_new_worker=true  }
    }
    if create_new_worker == true { fmt.Printf("\n<< Ask to create new worker >>\n")  ;  w := wp.AddWorker() ; w.Append(tgt) }

}
