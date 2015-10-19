//package chase
package main

import "time"


var FILES_PER_GR                           = 1000
var TIMEOUT_MS              time.Duration  = 200
var LOG_CHANNEL_TIMEOUT_MS  time.Duration  = 1000





type TaskManager struct {

    FunListFile *[](tgt *Target)     func(   )
    FunListFile *[](tgt *TargetDir)  func(   )


}

func ( tm *TaskManager ) Start ()  {


}

func ( tm *TaskManager ) Append (   ) {


}

