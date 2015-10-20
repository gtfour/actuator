package chase
//package main

import "time"
//import "fmt"


var FILES_PER_GR                           = 1000
var TIMEOUT_MS              time.Duration  = 200
var LOG_CHANNEL_TIMEOUT_MS  time.Duration  = 1000


type AbstractTarget interface {

    Chasing() error

}



type WorkerPool struct {

    Workers  []Worker

}


type Worker struct {

    //FunListFile *[](tgt *Target)     func(   )
    //FunListFile *[](tgt *TargetDir)  func(   )
    Targets     []*Target
    //TargetDirs  []*TargetDir


}

func ( w *Worker ) Start ( )  {

    for {
        for tgt := range w.Targets {

            w.Targets[tgt].Chasing()

        }

    }

}

func ( w *Worker ) Append ( tgt *Target ) {

    w.Targets = append(w.Targets,tgt)

}

func WPcreate () (wp WorkerPool, err error) {

    return wp, nil

}

func (wp *WorkerPool)  WPappend () {



}

func ( wp *WorkerPool ) AppendTarget ( tgt *Target ) () {


    var create_new_worker bool

    for w:= range wp.Workers {

        worker := wp.Workers[w]

        for wtgt := range worker.Targets {

            worker_target_dir := worker.Targets[wtgt]

            if tgt.Dir == worker_target_dir.Path { create_new_worker = true }

       }
       if create_new_worker == false { worker.Append(tgt)  }
    }
    //fmt.Printf("%t",create_new_worker)

}
