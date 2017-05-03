package filtering
//
import "errors"
//
var dup_name                = errors.New("error:filter with following name is already exist")
var name_is_none            = errors.New("error:filter name wasn't specified")
var offset_out_of_range     = errors.New("error:offset out of range")
var input_size_out_of_range = errors.New("error:input size with current offset is out of range")
//
