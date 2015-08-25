package main

import "fmt"
import "os"




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


    var hostname_providers =  map[string][]string  {"direct":{"/proc/sys/kernel/hostname","/etc/hostname","/etc/HOSTNAME"}, "complex":{"/etc/sysconfig/network"}}
    var complex_key_phrases = []string {"HOSTNAME"}

    return nil


}

func (os *OS) Name() (err error) {


    var name_providers = map[string][]string {"complex":{"/etc/SuSE-release", "/etc/lsb-release", "/etc/os-release"},"direct":"/etc/redhat-release", "/etc/fedora-release"}

    var complex_key_phrases = []string {"NAME"}

    return nil


}

func (os *OS) Version() (err error) {

    var name_providers = map[string][]string {"complex":{"/etc/SuSE-release", "/etc/lsb-release", "/etc/os-release"},"direct":"/etc/redhat-release","/etc/issue"}

    var complex_key_phrases = []string {"VERSION_ID","VERSION"}

    return nil


}

//

func (os *OS) Release() (err error) {

    var name_providers = map[string][]string {"complex":{"/etc/SuSE-release", "/etc/lsb-release", "/etc/os-release"},"direct":"/etc/redhat-release","/etc/issue"}

    var complex_key_phrases = []string {"VERSION_ID","VERSION","release"}

    return nil


}




func main() {

   operating_system:=&OS{OsName:"ubuntu",Files:[]string{"/etc/release"}}

   fmt.Println("----")

   fmt.Println(operating_system)

   fmt.Println("----")

}

func ReadFile(filename string, search_phrases []string) (params []string,err error) {

    var delimiters = []string {"="," ",": "}

    file, err := os.Open(filename)

    status_file:=StatusFile{}

    if err!=nil {

        return status_file,err

    }

    buffered_reader:=bufio.NewReader(file)
    eof := false

    status_entry:=StatusEntry{}

    for lino := 1; !eof; lino++ {


        line, err := buffered_reader.ReadString('\n')

        if err == io.EOF {
            err = nil
            eof = true
        } else if err != nil {
            return status_file, err
        }

        if ( strings.HasPrefix(line, "Package") || strings.HasPrefix(line, "Status") ||  strings.HasPrefix(line, "Architecture") || strings.HasPrefix(line, "Version")){
            status_entry.ParseField(line)
        }
        if (status_entry.Complete) && (status_entry.Installed) {

          status_file.InstalledPackages=append(status_file.InstalledPackages,status_entry)
          status_entry=StatusEntry{}

        }

    }

    return status_file,nil


}

