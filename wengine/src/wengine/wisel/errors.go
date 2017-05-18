package wisel

import "errors"

var PARAMS_NOT_ENOUGH  = errors.New("Not enough params to create new dashboard source (dynima).")
var CANT_FIND_WSCLIENT = errors.New("Can't find ws client.")
