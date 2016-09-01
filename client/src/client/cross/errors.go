package cross

import "errors"

var unable_to_open_db     = errors.New("\nErr:Unable to open database")
var unable_to_init        = errors.New("\nErr:Unable to init default collections")
/*
var dynimas_open_error    = errors.New("Unable to open dynimas collection")
var dynima_edit_error     = errors.New("Unable to edit dynima")
var dynima_get_error      = errors.New("Unable to get dynima")
var dynima_remove_error   = errors.New("Unable to remove dynima")
var dynimas_list_is_empty = errors.New("Dynimas list is empty")
var motion_write_error = errors.New("Unable to write motion")
*/

var collection_open_error          = errors.New("\nErr:Unable to open collection")
var collection_entry_edit_error    = errors.New("\nErr:Unable to edit collection entry")
var collection_entry_get_error     = errors.New("\nErr:Unable to get collection entry")
var collection_entry_remove_error  = errors.New("\nErr:Unable to remove collection entry")
var collection_entry_write_error   = errors.New("\nErr:Unable to write collection entry")
var collection_entry_list_is_empty = errors.New("\nErr:Collection entry list is empty")

var trigger_wasnt_assigned         = errors.New("\nErr:Trigger wasnt assigned to this target")

var table_doesnt_exist = errors.New("\nTable does'nt exist")




