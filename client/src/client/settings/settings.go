package settings

var RESTAPI_SETDASHBOARDDATA_URL = "/rest/dashboard/set-dashboard-data/"

var RESTAPI_SERVER_ADDR          = "10.10.111.143"
var RESTAPI_SERVER_PROTO         = "http"
var RESTAPI_SERVER_PORT          = "9000"
var RESTAPI_URL                  = RESTAPI_SERVER_PROTO+"://"+RESTAPI_SERVER_ADDR+":"+RESTAPI_SERVER_PORT

var RESTAPI_WS_DATA_URL          = "/entry"
var RESTAPI_WS_PROTO             = "ws"

var RESTAPI_WS_URL               = RESTAPI_WS_PROTO+"://"+RESTAPI_SERVER_ADDR+RESTAPI_WS_DATA_URL
var RESTAPI_WS_ORIGIN            = RESTAPI_SERVER_PROTO+"://"+RESTAPI_SERVER_ADDR

var SERVER_ID                    = ""
var CLIENT_ID                    = ""

//var PKI_DIR                      = "/top_secret_dir"
var PUBKEY_PATH                  = "/top_secret_dir/pub.rsa"
var PRIVATEKEY_PATH              = "/top_secret_dir/private.rsa"

