package riva

//import "encoding/json"
import "client/activa"
//import "client/evebridge"

func Handle()(error){

    return nil


}



type Trigger interface {

    SetMeow(*Meow)
    GetMeow(meow_id string)(*Meow)
    SetFractal(*Fractal)
    GetFractal(fractal_id string)(*Fractal)
    Remove()

}

type Meow interface {

    Set()
    Write()
    Remove()

}

type Fractal interface {

    Set()
    Run()
    Remove()

}

type Runner struct {

}

type Editor struct {

}

func CreateTrigger()(t *Trigger,err error) {



    return
}

func MakeShiver(t *Trigger)(*activa.Motion){

    var m activa.Motion
    m.TaskState = activa.TASK_STATE_new


    return &m
}
