//package parse

package main

//import "fmt"

type Config struct {

    start bool
    print_help bool
    request_config_from_wengine bool
    wengine_address string

}

var wengine_server_address_xmltag = "wengine-ip"

var config_file_name = "actuator.conf"

//


var words = map[string][]string{"start": {"-s","start","--start"},
                                "wengine": {"-w","wengine-address","--wengine-address"},
                                "request_config":{"-r","request-config","--request-config"},
                                "print_help":{"-h","help","--help"} }

//








func Parse(osargs []string) (config *Config) {


    config=&Config{}


   for num:= range osargs {

   




   }


    return config


}

func Do() {


}

func Search (word_list []string, word string )(bool) {

    marker:=false

    for num:=range word_list {

        if word_list[num] == word {

            marker=true

        }

    }

    return marker

}

func main() {


}

