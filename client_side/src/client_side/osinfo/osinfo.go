package main

import "fmt"
import "os"
import "bufio"




type OS struct {

    HostName string
    OsName string
    OsVersion string
    OsRelease string
    Files []string

}

func (os *OS) CheckEtcFiles() (err error) {

    fmt.Println(os.Files)

    return nil

}


func (os *OS) Hostname() (err error) {


    var providers =  map[string][]string  {"direct":{"/proc/sys/kernel/hostname","/etc/hostname","/etc/HOSTNAME"}, "complex":{"/etc/sysconfig/network"}}
    var complex_keys = []string {"HOSTNAME"}

    return nil


}

func (os *OS) Name() (err error) {


    var providers = map[string][]string {"complex":{"/etc/SuSE-release","/etc/SuSE-brand" , "/etc/lsb-release", "/etc/os-release"},"direct":{"/etc/redhat-release", "/etc/fedora-release","/etc/SuSE-brand"}}

    var complex_keys = []string {"NAME"}

    return nil


}

func (os *OS) Version() (err error) {

    var providers = map[string][]string {"complex":{"/etc/SuSE-release", "/etc/lsb-release", "/etc/os-release", "/etc/SuSE-brand"},"direct":{"/etc/redhat-release","/etc/issue"}}

    var complex_keys = []string {"VERSION_ID","VERSION","release"}

    return nil


}

//

func (os *OS) Release() (err error) {

    var providers = map[string][]string {"complex":{"/etc/lsb-release","/etc/fedora-release","/etc/redhat-release"}}

    var complex_keys = []string {"release","DISTRIB_RELEASE"}

    return nil


}




func main() {

   operating_system:=&OS{OsName:"ubuntu",Files:[]string{"/etc/release"}}

   fmt.Println("----")

   fmt.Println(operating_system)

   fmt.Println("----")

}

func ReadFile(providers map[string][]string , complex_keys []string) (value string,err error) {

    var delimiters = []string {"="," ",": "}
    value = "Unknown"
    for provider_type := range providers {

        if provider_type == "direct" {

            for name := range providers["direct"] {

                filename := providers["direct"][name]
                file, err := os.Open(filename)
                if err==nil {
                     buffered_reader:=bufio.NewReader(file)

                }

        }} else {



       }


    }
    return "",nil


}
