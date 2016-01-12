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

    Name      string
    Type      int8
    Path      string
    Pool      Pool
    Directory string
    IsDir     bool


}

type Pool struct {

    Name     string
    Actions  []Action
    Files    []File
    Feature  []Feature

}

type Trigger struct {

}
