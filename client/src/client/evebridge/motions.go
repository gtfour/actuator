package evebridge

import "client/cross"
import "client/activa"

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

