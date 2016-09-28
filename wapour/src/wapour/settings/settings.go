package settings
import "html/template"
import "github.com/gin-gonic/gin"
//##################################
//##################################
//##################################
//const SERVER_ADDR          = "127.0.0.1"
const SERVER_ADDR string  = "10.10.111.143"
const SERVER_PROTO         = "http"
const SERVER_PORT          = "8090"
const SERVER_URL           = SERVER_PROTO+"://"+SERVER_ADDR+":"+SERVER_PORT
//const RESTAPI_SERVER_ADDR  = "127.0.0.1"
const RESTAPI_SERVER_ADDR  = "10.10.111.143"
const RESTAPI_SERVER_PROTO = "http"
const RESTAPI_SERVER_PORT  = "9000"
const WS_LINE              = "/entry"
const RESTAPI_URL          = RESTAPI_SERVER_PROTO+"://"+RESTAPI_SERVER_ADDR+":"+RESTAPI_SERVER_PORT
const HTTP_WS_URL          = SERVER_URL+"/entry"
const WS_URL               = template.URL("ws://"+SERVER_ADDR+":"+SERVER_PORT+WS_LINE)
//const WS_URL               = template.URL("ws://127.0.0.1:8090/entry")
const GET_DATA_URI         = "/dashboard/get-dashboard-data/"
const GET_DATA_URL         = SERVER_URL+GET_DATA_URI
//const GET_DATA_URL         = "http://127.0.0.1:8090/userspace/get-data"
//##################################
//##################################
//##################################
const TOKEN_COOKIE_FIELD_NAME  string = "USER_TOKEN"
const USERID_COOKIE_FIELD_NAME string = "USER_ID"
//##################################
//##################################
//##################################
const STATIC_DIR                = "/actuator/wapour/static"
const STATIC_URL                = "/static/main/"
//##################################
//##################################
//##################################
const ONLINE_USERS_DB_FILE      = "/tmp/users.db"
//const ONLINE_USERS_STORAGE_TYPE = "db" // "ram"
const ONLINE_USERS_STORAGE_TYPE = "ram"
//##################################
//##################################
//##################################
const APP_SETTINGS       = gin.H{ "static_url":STATIC_URL, "ws_url":WS_URL, "get_data_url":GET_DATA_URL }
const USERSPACE_DATA_URL ="/userspace-data"
const ADMIN_DATA_URL     ="/index-data"
//##################################
//##################################
//##################################
const ALLOWED_REDIRECTS = []string { "/userspace","/index" }
