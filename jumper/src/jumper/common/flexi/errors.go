package flexi

import "errors"

var notStringSlice     = errors.New("flexi:input interface is not a string slice")
var notInterfaceSlice  = errors.New("flexi:input interface is not a interface slice")
var notInt             = errors.New("flexi:input interface has not int type ")
var notIntSlice        = errors.New("flexi:input interface has not int-slice type ")
var indexOutOfRange    = errors.New("flexi:index out of range")
