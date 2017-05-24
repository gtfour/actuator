package evebridge

import "fmt"
import "client/cross"
// import "client/activa"
import "jumper/activa"


func HandleMotion(motion *activa.Motion)(err error) {
    if motion.TaskState == activa.TASK_STATE_new || motion.TaskState == activa.TASK_STATE_empty {
        err=HandleNewMotion(motion)
        return err
    }
    return unrecognizable_task_state
}

func HandleNewMotion(motion *activa.Motion)(err error) {
    err=cross.WriteMotion(motion)
    if err == nil {
        motion.TaskState=activa.TASK_STATE_inprogress
    } else {
        motion.TaskState=activa.TASK_STATE_failed
        return err
    }
    return err
}

func HandleCompletedMotion(motion *activa.Motion)(err error) {
    _=cross.WriteMotion(motion)
    return err
}




func (a *App)handleMotion(motion *activa.Motion)(){

    //
    // fmt.Printf("\nNew motion %v\n", motion)
    //
    motion.TaskState =  activa.TASK_STATE_inprogress
    motionSubType    := motion.SubType
    motionSourceType := motion.SourceType
    motionSourcePath := motion.SourcePath
    _,_              =  motionSourceType, motionSourcePath
    // // cross.WriteMotion( &motion )
    //
    //
    switch motion_type := motion.Type; motion_type {
        //
        //
        case activa.MOTION_TYPE_BLACKOUT:
            //
            // // commands
            //
            //
        case activa.MOTION_TYPE_BLACKTOP:
            //
            // // files and directories
            if motionSubType == activa.MOTION_SUBTYPE_ADD_DYNIMA {
                fmt.Printf("\nMotion %v  Data: motionSubType: %v  motionSubType: %v motionSourceType: %v motionSourcePath: %v \n", motion_type, motionSubType, motionSourceType, motionSourcePath )
                }
                //
                //
            }
            //
            //
            //
            a.motions<-motion
}
