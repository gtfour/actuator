//package parse

package cmdlineparse

type Config struct {

    start                       bool
    print_help                  bool
    request_config_from_wengine bool
    wengine_address             string
    mistake                     bool

}

var wengine_server_address_xmltag = "wengine-ip"

var config_file_name              = "actuator.conf"

//


var words = map[string][]string{"start": {"-s","start","--start"},
                                "wengine_address": {"-w","wengine-address","--wengine-address"},
                                "request_config":{"-r","request-config","--request-config"},
                                "print_help":{"-h","help","--help"} }

func Parse(osargs []string) (config Config) {


    config=Config{}

   for osarg_num    :=  range osargs {

       osarg_word   :=  osargs[osarg_num]

       for word_key :=  range words {

           found:=Search(words[word_key],osarg_word)

           if found {

               if word_key == "start" { config.start=true }
               if word_key == "wengine_address" { config.wengine_address=osarg_word }
               if word_key == "request_config" { config.request_config_from_wengine=true }
               if word_key == "print_help" { config.print_help=true }

           }
       }
   }

   return config

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
