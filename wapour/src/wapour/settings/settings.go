package settings
import "html/template"
//
//var SERVER_ADDR          = "127.0.0.1"
var SERVER_ADDR          = "10.10.111.143"
var SERVER_PROTO         = "http"
var SERVER_PORT          = "8090"
var SERVER_URL           = SERVER_PROTO+"://"+SERVER_ADDR+":"+SERVER_PORT
//var RESTAPI_SERVER_ADDR  = "127.0.0.1"
var RESTAPI_SERVER_ADDR  = "10.10.111.143"
var RESTAPI_SERVER_PROTO = "http"
var RESTAPI_SERVER_PORT  = "9000"
var RESTAPI_URL          = RESTAPI_SERVER_PROTO+"://"+RESTAPI_SERVER_ADDR+":"+RESTAPI_SERVER_PORT
var WS_URL               = template.URL("ws://10.10.111.143:8090/entry")
var GET_DATA_URL         = "http://10.10.111.143:8090/userspace/get-data"
//
//
var TOKEN_COOKIE_FIELD_NAME  string = "USER_TOKEN"
var USERID_COOKIE_FIELD_NAME string = "USER_ID"
//
//
var STATIC_DIR                = "/actuator/wapour/static"
var STATIC_URL                = "/static/main/"
var ONLINE_USERS_DB_FILE      = "/tmp/users.db"
//var ONLINE_USERS_STORAGE_TYPE = "db" // "ram"
var ONLINE_USERS_STORAGE_TYPE = "ram"
