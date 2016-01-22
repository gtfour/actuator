package wengine


type _Seed struct {

    Name    string
    Action  _Action
    Value   string

}

type _Feature struct {

    Name   string
    Seeds  []_Seed

}

type _Action struct {

    Name    string
    Command string

}

type _File struct {

    Name      string
    Type      int8
    Path      string
    Pool      _Pool
    Directory string
    IsDir     bool


}

type _Pool struct {

    Name     string
    Actions  []_Action
    Files    []_File
    Feature  []_Feature

}

type _Trigger struct {

}

type _Host struct {

    Id       string
    Name     string
    Features []_Feature

}
