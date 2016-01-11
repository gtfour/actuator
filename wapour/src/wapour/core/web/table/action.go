package table

var ATYPE_MODAL      int8 = 0
var ATYPE_LINK       int8 = 1

var MODALTYPE_REMOVE int8 = 2
var MODALTYPE_CREATE int8 = 3
var MODALTYPE_EDIT   int8 = 4


type Action struct {

    Name        string
    Type        int8
    SuccessUrl  string
    PostUrl     string
    RedirectUrl string

}

type ModalWindow struct {

    ModalName        string
    ModalId          string
    ModalTitle       string
    ModalHelpMessage string
    FieldSet         []Field

}

type Field struct {

    Name         string
    Type         string
    Value        string
    DefaultValue string
    Required     bool

}

