package petrovich

import "fmt"
import "jumper/common/file"

func LoadConfig(initial_config map[string]string)(config map[string]string,err error){

    /*lines , err:=*/_,err=file.ReadFile(filename)
    if err!= nil {
        return nil,err
    }
    var parser huyamba
    if len(initial_config)>0{
        parser=CreateHuyamba(initial_config[0])
    } else {
        parser=CreateHuyamba()
    }
    fmt.Printf("Parser:\n%v\n",parser)
    return config, nil

}

func ExtractBaseSet(initial_config map[string]string)(config_path string, config_type string, err error){
    return config_path, config_type, err
}
