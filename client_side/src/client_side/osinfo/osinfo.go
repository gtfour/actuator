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
    var name_providers = map[string][]string {"complex":{"/etc/redhat-release","/etc/lsb-release"}}

    return nil

}

func (os *OS) Hostname() (err error) {

    var name_providers = map[string][]string {"complex":{"/etc/redhat-release","/etc/lsb-release"}}


}



func main() {

   operating_system:=&OS{OsName:"ubuntu",Files:[]string{"/etc/release"}}

   fmt.Println("----")

   fmt.Println(operating_system)

   fmt.Println("----")

}

func ReadFile(filename string, search_phrases []string) (params []string,err error) {

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

