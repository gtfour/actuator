package config

var wengine_server_address_xmltag = "wengine-ip"

var config_file_name = "actuator.conf"

type ConfigFile struct {


}

func CreateConfigFile()(err error, config_file *ConfigFile) {

    return nil, &ConfigFile{}


}

func (configfile *ConfigFile) GetSelfPath() (err error){

    return nil

}

func (configfile *ConfigFile) AddDataToBlock(block_name string) (err error){

    return nil


}

func (configfile *ConfigFile) RequestNewConfig(block_name string) (err error){

    return nil

}


