package run

var default_settings = getDefaultSettings()

func getDefaultSettings()(settings map[string]string){

    settings                         = make(map[string]string)
    settings["SERVER_ADDR"]          = "0.0.0.0"
    settings["SERVER_PROTO"]         = "0.0.0.0"
    settings["SERVER_PORT"]          = "80"
    settings["SERVER_URL"]           = settings["SERVER_PROTO"]+"://"+settings["SERVER_ADDR"]+settings["SERVER_PORT"]

    settings["RESTAPI_SERVER_ADDR"]  = "127.0.0.1"
    settings["RESTAPI_SERVER_PROTO"] = "http"
    settings["RESTAPI_SERVER_PORT"]  = "9000"
    settings["RESTAPI_URL"]          = settings["RESTAPI_SERVER_PROTO"] +"://"+settings["RESTAPI_SERVER_ADDR"]+":"+settings["RESTAPI_SERVER_PORT"]
    return

}

func GetDefaultPropValue(key string)(value string, err error){
     return
}
