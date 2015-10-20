package chase
//package main

import "time"
//import "fmt"


var FILES_PER_GR                           = 1000 // if FILES_PER_GR is very big - TargetsCount type should be modified 
var TIMEOUT_MS              time.Duration  = 200
var LOG_CHANNEL_TIMEOUT_MS  time.Duration  = 1000


type AbstractTarget interface {

    GetDir()  string
    Chasing() error
    //GetType() string
    GetPath() string

}

//func (a *AbstractTarget) GetDir() string { return a.Dir }



type WorkerPool struct {

    Workers         []*Worker
    WKillers        []chan bool

}


type Worker struct {

    //FunListFile *[](tgt *Target)     func(   )
    //FunListFile *[](tgt *TargetDir)  func(   )
    Targets       []AbstractTarget
    TargetsCount  int32
    Stop          chan bool
    //TargetDirs  []*TargetDir


}

func ( w *Worker ) Start ()  {

    for {
        select {

            case <-w.Stop:

                return

            default:

                for tgt := range w.Targets {
                    w.Targets[tgt].Chasing()
                }

        }

    }

}

func ( w *Worker ) Append ( tgt AbstractTarget ) {

    //if tgt.GetType() == "dir" {

        w.Targets = append(w.Targets,tgt)

    //} else {

        //w.Targets = append(w.Targets,tgt.(*Target))

    //}

}

func WPCreate () (wp WorkerPool) {

    wp.AddWorker()

    workers := wp.Workers

    for i:= range workers {

        go workers[i].Start()

    }
    return wp

}

func ( wp *WorkerPool ) Stop () {

    killers := wp.WKillers

    for i:= range killers {

        killers[i] <- true

    }

}


func (wp *WorkerPool)  AddWorker()(w *Worker){

    //w           =   &Worker{}
    w.Stop      =   make(chan bool)
    wp.Workers  =   append(wp.Workers, w)
    return


}

func ( wp *WorkerPool ) AppendTarget ( tgt AbstractTarget ) () {

    var create_new_worker bool

    for w:= range wp.Workers {

        worker := wp.Workers[w]

        for wtgt := range worker.Targets {

            worker_target_dir := worker.Targets[wtgt]

            if tgt.GetDir() == worker_target_dir.GetPath() { create_new_worker = true }

       }
       if create_new_worker == false { worker.Append(tgt) ; break  }
    }
    if create_new_worker == true { w := wp.AddWorker() ; w.Append(tgt) }
    //fmt.Printf("%t",create_new_worker)

}
