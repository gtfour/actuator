package main

import (
        "os"
        "fmt"
        //"encoding/xml"
)

func GetConfFile()(conffile string ,err error) {

    if (len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help"))||(len(os.Args) == 1) {
        err = fmt.Errorf("usage: %s --conf || -c conffile.conf", filepath.Base(os.Args[0]))
        return "", err
    }
    
    if len(os.Args) > 1 {
        conffile = os.Args[1]
    }

    return conffile, err


}





func main() {

   var conf_file string 
  
   conf_file = "actuator.conf"

   xmlFile, err := os.Open("actuator.conf")

   if err != nil {

       fmt.Println("Error opening file",conf_file)
       return 
    }
    defer xmlFile.Close() 
    //xml.Unmarshal(xmlFile, &q)
}
