//package config
package main
import ( "os" ; "path/filepath" ; "fmt" ; "io/ioutil"  )

var config_delim = "|"
var repositoryTag = "[repository]"
var packageTag = "[package]"
var selfconfTag="[selfconf]"
var wengineAddressTag = "[wengine]"
var configFileName = "actuator.conf"

func main() {

  config:=&Config{}
  fmt.Println(config.GetCurrentConfigFile())

}

type Config struct {

  ConfigFilePath string
  Files []string
  ConfigFileExists bool
  

}

func CreateConfigFile()(err error, cnf *Config) {

    return nil,cnf

}

func (cnf *Config) GetCurrentConfigFile() (err error){

    dir, _ := filepath.Split(os.Args[0])
    fmt.Printf("\nDirectory:  %s \n",dir)
    if _, err := os.Open(dir+configFileName) ; err != nil {

        cnf.ConfigFileExists=false
        return err

    } else {

        err=cnf.RequestNewConfig()
        if err!=nil {

            return err

        }
    }
    cnf.ConfigFileExists=true

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

func (cnf  *Config)ParseFile (err error) {

    if (cnf.ConfigFileExists) {

      content, err := ioutil.ReadFile(cnf.ConfigFilePath)
      for line_num :=range content {

      }


}
