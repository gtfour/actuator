package run

var default_props = getDefaultSettings()

func getDefaultSettings()(prop map[string]string){

    prop                         = make(map[string]string)
    prop["SERVER_ADDR"]          = "0.0.0.0"
    prop["SERVER_PROTO"]         = "0.0.0.0"
    prop["SERVER_PORT"]          = "80"
    prop["SERVER_URL"]           = prop["SERVER_PROTO"]+"://"+prop["SERVER_ADDR"]+prop["SERVER_PORT"]

    prop["RESTAPI_SERVER_ADDR"]  = "127.0.0.1"
    prop["RESTAPI_SERVER_PROTO"] = "http"
    prop["RESTAPI_SERVER_PORT"]  = "9000"
    prop["RESTAPI_URL"]          = prop["RESTAPI_SERVER_PROTO"] +"://"+prop["RESTAPI_SERVER_ADDR"]+":"+prop["RESTAPI_SERVER_PORT"]

    return

}

func GetDefaultPropValue(key string)(value string, err error){
     return
}
