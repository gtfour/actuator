package main

import "jumper/petrovich"

func main(){

    var appConfig petrovich.Huyamba
    appConfig.Description   = "testing Petrovich functionality"
    appConfig.AppName       = "test_petrovich"


    initial_config                := make(map[string]string, 0)
    initial_config["config_path"] =  "/etc/wengine/wengine.conf"
    initial_config["config_type"] =  "file"
    petrovich.ExtractBaseSet(initial_config)

}
