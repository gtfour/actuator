package cross

import "errors"

var unable_to_init        = errors.New("Unable to init default collections")
var dynimas_open_error    = errors.New("Unable to open dynimas collection")
var dynima_edit_error     = errors.New("Unable to edit dynima")
var dynima_get_error      = errors.New("Unable to get dynima")
var dynima_remove_error   = errors.New("Unable to remove dynima")
var dynimas_list_is_empty = errors.New("Dynimas list is empty")

