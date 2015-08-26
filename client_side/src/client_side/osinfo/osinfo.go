package main

import "fmt"
import "os"
import "bufio"
import "io"
// I am glad to introduce you "f-shit power" and "How does this fking code work" technologies 



type OS struct {

    Hostname string
    OsName string
    OsVersion string
    OsRelease string
    Files []string
    VirtualProvider []string // по задумке сюда будут складываться строки начиная со второй, которые были найдены
                             // проходом значений хеша по ключу директ 
                             // по идее директ должен включать файлы имеющие всего одну строку, не содержащую ключей а только одно значение
                             // прошу прощения за этот бред )

}



func main() {

   operating_system:=&OS{}

   operating_system.GetHostname()
   fmt.Println(operating_system.Hostname)
   fmt.Println(operating_system.VirtualProvider)

   fmt.Println("----")

}



func (os *OS) GetHostname() (err error) {

                                          // по ключу директ находятся однострочные файлы
                                          // по ключу комплекс - файлы которые надо парсить на предмет наличия внутри complex_key
    var providers =  map[string][]string  {"direct":{"/proc/sys/kernel/hostname","/etc/hostname","/etc/HOSTNAME"}, "complex":{"/etc/sysconfig/network"}}
    var complex_keys = []string {"HOSTNAME"}

    //fmt.Printf("\n:debug:\n%s\n%s",providers,complex_keys)

    os.Hostname,_,os.VirtualProvider=GetParamValue(providers,complex_keys)

    return nil

}

func (os *OS) Name() (err error) {


    var providers = map[string][]string {"complex":{"/etc/SuSE-release","/etc/SuSE-brand" , "/etc/lsb-release", "/etc/os-release"},"direct":{"/etc/redhat-release", "/etc/fedora-release","/etc/SuSE-brand"}}

    var complex_keys = []string {"NAME"}

    fmt.Printf("\n:debug:\n%s\n%s",providers,complex_keys)
    return nil



}

func (os *OS) Version() (err error) {

    var providers = map[string][]string {"complex":{"/etc/SuSE-release", "/etc/lsb-release", "/etc/os-release", "/etc/SuSE-brand"},"direct":{"/etc/redhat-release","/etc/issue"}}

    var complex_keys = []string {"VERSION_ID","VERSION","release"}

    fmt.Printf("\n:debug:\n%s\n%s",providers,complex_keys)
    return nil

}

//

func (os *OS) Release() (err error) {

    var providers = map[string][]string {"complex":{"/etc/lsb-release","/etc/fedora-release","/etc/redhat-release"}}

    var complex_keys = []string {"release","DISTRIB_RELEASE"}

    fmt.Printf("\n:debug:\n%s\n%s",providers,complex_keys)
    return nil


}



func GetParamValue(providers map[string][]string , complex_keys []string) (value string,err error,vp []string) {

    //    var delimiters = []string {"="," ",": "}

    value = "Unknown"

    var possible_value_candidates []string

    for provider_type := range providers {

        if provider_type == "direct" {

            for name := range providers["direct"] {

                filename := providers["direct"][name]

                lines,err:=ReadFileLines(filename)

                

                if err==nil {

                     fmt.Println(len(lines))

                     lines_len :=len(lines)

                     if lines_len > 0 { possible_value_candidates=append(possible_value_candidates,lines[0]) }

                     if lines_len > 1 { for i:= range lines { vp=append(vp,lines[i]) } }

                }

        }} else {



       }


    }
    return "",nil,vp


}


func ReadFileLines (filename string) (lines []string,err error){

    file, err := os.Open(filename)

    if err!=nil {

        return lines, err

    }

    buffered_reader:=bufio.NewReader(file)

    eof := false

    for lino := 1; !eof; lino++ {

        line, err := buffered_reader.ReadString('\n')

        if err == io.EOF {

            err = nil

            eof = true

        } else if err != nil {

            return lines, err

        }

        lines=append(lines,line)

    }
    return lines,err

}

