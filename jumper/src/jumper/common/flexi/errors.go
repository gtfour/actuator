package flexi

import "errors"

var notStringSlice     = errors.New("flexi:input interface is not a string slice")
var notInterfaceSlice  = errors.New("flexi:input interface is not a interface slice")
