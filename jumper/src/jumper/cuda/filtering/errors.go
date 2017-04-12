package filtering

import "errors"

var dup_name         =  errors.New("error:filter with following name is already exist")
var name_is_none     =  errors.New("error:filter name wasn't specified")
