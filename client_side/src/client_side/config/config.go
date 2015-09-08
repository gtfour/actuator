package config
import ( "os" ; "path/filepath" )

var config_delim = "|"
var repositoryTag = "[repository]"
var packageTag = "[package]"
var selfconfTag="[selfconf]"
var wengineAddressTag = "[wengine]"
var configFileName = "actuator.conf"

type Config struct {

  Path string
  Files []string
  

}

func CreateConfigFile()(err error, cnf *Config) {

    return nil, &ConfigFile{}


}

func (cnf *Config) GetCurrentConfigFile() (err error){


    dir, _ := filepath.Split(os.Args[0])
    if configfile, err := os.Open(dir+configFileName) ; 
     
    return nil

}

func (cnf *Config) AddDataToBlock(blockTag string) (err error){

    return nil


}

func (cnf *Config) RequestNewConfig(block_name string) (err error){

    return nil

}

func (cnf  *Config)SetConfigPath(path string)(err error) {

  cnf.Path=path
  return nil


}

func (cnf  *Config)UploadToFile (err error) {





}
