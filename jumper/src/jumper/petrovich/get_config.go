package petrovich

import "fmt"
//import "jumper/common/file"

func LoadConfig(initial_config map[string]string)(config map[string]string,err error){

    /*lines , err:=*///_,err=file.ReadFile(filename)
    //if err!= nil {
   //     return nil,err
    //}
    parser:=CreateHuyamba(initial_config)
    fmt.Printf("Parser:\n%v\n",parser)
    return config, nil

}

func ExtractBaseSet(initial_config map[string]string)(config_path string, config_type string, err error){


    config_path,ok_config_path:=initial_config["config_path"]
    config_type,ok_config_type:=initial_config["config_type"]
    fmt.Printf("Config Path:%s\nConfig Type:%s\nConfig Path Extract Err:%v\nConfig Type Extract Err:%v\n",config_path,config_type,ok_config_path,ok_config_type)

    return config_path, config_type, err
}
