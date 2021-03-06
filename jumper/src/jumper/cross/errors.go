package cross
import "errors"

var db_type_is_incorrect = errors.New("\ncross:wrong db type")
var db_username_is_empty = errors.New("\ncross:username is empty")
var db_password_is_empty = errors.New("\ncross:password is empty")
var db_path_is_empty     = errors.New("\ncross:path is empty")
var db_dbname_is_empty   = errors.New("\ncross:dbname is empty")
var db_host_is_empty     = errors.New("\ncross:host is empty")

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
var EmptyTableName      = errors.New("cross:table name is empty")
var KeyAndValueEmpty    = errors.New("cross:key and Value are empty")
var KeyIsEmpty          = errors.New("cross:key is empty")
var ValueIsEmpty        = errors.New("cross:value is empty")
var IncorrectQueryType  = errors.New("cross:incorrect query type")
//
var SliceNameIsEmpty        = errors.New("cross:slice name is empty")
var SliceDoesntExist        = errors.New("cross:slice doesn't exist")
var NothingIsAppendToSlice  = errors.New("cross:nothing Is Append To Slice")
var RemoveIndexIsEmpty      = errors.New("cross:nothing to remove from slice. Remove index is empty")
var EntryIdIsEmpty          = errors.New("cross:entry id is empty")
//
