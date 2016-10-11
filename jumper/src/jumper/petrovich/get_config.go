package petrovich

import "jumper/common/file"

func GetConfigFromFile(filename string)(config map[string]string,err error){
    lines,err:=file.ReadFile(filename)
    if err!= nil {
        return nil,err
    }
    return config, nil
}
