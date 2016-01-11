package wengine


type Seed struct {

    Name    string
    Action  Action
    Value   string

}

type Feature struct {

    Name   string
    Seeds  []Seed

}

type Action struct {

    Name    string
    Command string

}

type File struct {


}

type Trigger struct {

}
