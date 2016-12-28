package cross
import "errors"

var db_type_is_incorrect = errors.New("\ncross:wrong db type")
var db_username_is_empty = errors.New("\ncross:username is empty")
var db_password_is_empty = errors.New("\ncross:password is empty")
var db_path_is_empty     = errors.New("\ncross:path is empty")
var db_dbname_is_empty   = errors.New("\ncross:dbname is empty")

var cant_open_database   = errors.New("\ncross:Can't open database")

var Selected_dbtype_is_not_ok_on_client_side = errors.New("\ncross:selected dbtype is not appropriate for using on client side")
var Selected_dbtype_is_not_ok_on_server_side = errors.New("\ncross:selected dbtype is not appropriate for using on server side")


// key-value database errors

var TableDoesntExist   = errors.New("Table does'nt exist")
var EntryDoesntExist   = errors.New("Entry does'nt exist")
var EntryAlreadyExist  = errors.New("Entry is already exist")
var EncodeError        = errors.New("Encode error")
var DecodeError        = errors.New("Decode error")

// database query errors

var empty_key            = errors.New("Key is empty")
var empty_query          = errors.New("Query is empty")
var key_and_value_empty  = errors.New("Key and Value are empty")
var incorrect_query_type = errors.New("Incorrect query type")


