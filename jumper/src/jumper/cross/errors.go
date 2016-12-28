package cross
import "errors"

var db_type_is_incorrect = errors.New("\ncross:wrong db type")
var db_username_is_empty = errors.New("\ncross:username is empty")
var db_password_is_empty = errors.New("\ncross:password is empty")
var db_path_is_empty     = errors.New("\ncross:path is empty")
var db_dbname_is_empty   = errors.New("\ncross:dbname is empty")

var CantOpenDatabase   = errors.New("\ncross:Can't open database")

var Selected_dbtype_is_not_ok_on_client_side = errors.New("\ncross:selected dbtype is not appropriate for using on client side")
var Selected_dbtype_is_not_ok_on_server_side = errors.New("\ncross:selected dbtype is not appropriate for using on server side")


// key-value database errors

var TableDoesntExist   = errors.New("cross:table does'nt exist")
var EntryDoesntExist   = errors.New("cross:entry does'nt exist")
var EntryAlreadyExist  = errors.New("cross:entry is already exist")
var EncodeError        = errors.New("cross:encode error")
var DecodeError        = errors.New("cross:decode error")

// database query errors

var EmptyKey            = errors.New("cross:key is empty")
var EmptyQuery          = errors.New("cross:query is empty")
var KeyAndValueEmpty    = errors.New("cross:key and Value are empty")
var IncorrectQueryType  = errors.New("cross:incorrect query type")


