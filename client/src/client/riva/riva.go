package riva

import "client/activa"
//import "client/evebridge"

func Handle()(error){

    return nil


}



type Trigger interface {

    SetMeow(*Meow)
    GetMeow(meow_id string)
    SetFractal(*Fractal)
    GetFractal(fractal_id string)
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

func MakeShiver()(m *activa.Motion){


    return
}
