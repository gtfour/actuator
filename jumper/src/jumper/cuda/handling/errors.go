package handling

import "errors"

var filterListIsNil     =  errors.New("cuda:handling:Filter list could'nt be empty\n")
var targetIsNil         =  errors.New("cuda:handling:Target could'nt be empty\n")
var targetTypeUndefined =  errors.New("cuda:handling:Target type undefined\n")
var unableToHandleLine  =  errors.New("cuda:handling:line could'nt be parsed\n")

