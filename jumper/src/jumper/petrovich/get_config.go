package petrovich

import "jumper/common/file"

func LoadConfig(filename string, initial_config ...map[string]string)(config map[string]string,err error){
    lines,err:=file.ReadFile(filename)
    if err!= nil {
        return nil,err
    }
    var parser huyamba
    if len(initial_config)>0{
        parser=CreateHuyamba(initial_config[0])
    } else {
        parser=CreateHuyamba()
    }


    return config, nil
}
