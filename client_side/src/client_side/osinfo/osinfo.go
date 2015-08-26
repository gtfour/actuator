package main

import "fmt"
import "io/ioutil"
import "strings"
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


}



func (os *OS) GetHostname() (err error) {

//по ключу директ находятся однострочные файлы
//по ключу комплекс - файлы которые надо парсить на предмет наличия внутри complex_key


    var providers =  map[string][]string  {"direct":{"/proc/sys/kernel/hostname","/etc/hostname","/etc/HOSTNAME"}, "complex":{"/etc/sysconfig/network"}}
    var complex_keys = []string {"HOSTNAME"}

    key :="hostname"

    //fmt.Printf("\n:debug:\n%s\n%s",providers,complex_keys)

    var values []string

    values,_,os.VirtualProvider=GetParamValue(providers,complex_keys)
    os.Hostname,_=ValidateValue(values,key)

    return nil

}

func (os *OS) Name() (err error) {


    var providers = map[string][]string {"complex":{"/etc/SuSE-release","/etc/SuSE-brand" , "/etc/lsb-release", "/etc/os-release"},"direct":{"/etc/redhat-release", "/etc/fedora-release","/etc/SuSE-brand"}}

    var complex_keys = []string {"NAME"}

    key:="name"

    fmt.Printf("\n:debug:\n%s\n%s",providers,complex_keys)
    return nil



}

func (os *OS) Version() (err error) {

    var providers = map[string][]string {"complex":{"/etc/SuSE-release", "/etc/lsb-release", "/etc/os-release", "/etc/SuSE-brand"},"direct":{"/etc/redhat-release","/etc/issue"}}

    var complex_keys = []string {"VERSION_ID","VERSION","release"}

    key:="version"

    fmt.Printf("\n:debug:\n%s\n%s",providers,complex_keys)
    return nil

}

//

func (os *OS) Release() (err error) {

    var providers = map[string][]string {"complex":{"/etc/lsb-release","/etc/fedora-release","/etc/redhat-release"}}

    var complex_keys = []string {"release","DISTRIB_RELEASE"}

    key:="release"

    fmt.Printf("\n:debug:\n%s\n%s",providers,complex_keys)

    return nil


}



func GetParamValue(providers map[string][]string , complex_keys []string) (values []string,err error,vp []string) {


    var possible_value_candidates []string

    for provider_type := range providers {

        if provider_type == "direct" {

            for name := range providers["direct"] {

                filename := providers["direct"][name]

                lines,err:=ReadFileLines(filename)

                if err==nil {

                     lines_len :=len(lines)

                     if lines_len > 0 { possible_value_candidates=append(possible_value_candidates,lines[0]) }

                     if lines_len > 1 { for i:= range lines { if  lines[i] != "" {  vp=append(vp,lines[i]) }  }  }

                }

        }} else {

            for name := range providers["complex"] {

                filename := providers["complex"][name]

                lines,err:=ReadFileLines(filename)

                for ckey:= range complex_keys {

                    for num := range lines {

                        if strings.HasPrefix(lines[num], complex_keys[ckey]) {



                        }

                    }

                }

            }
       }


    }
    return possible_value_candidates,nil,vp


}


func ReadFileLines (filename string) (lines []string,err error){


    content, err := ioutil.ReadFile(filename)

    if err != nil {

        return lines, err

    }

    lines = strings.Split(string(content), "\n")

    return lines,err

}

func ValidateValue (values []string, key string) (value string,err error) {

    for i := range values{

       if key == "hostname" {

           if (len(values[i])>len(value))&&(! strings.HasPrefix(values[i], "local")) { value=values[i] }

       }


    }
    return value, nil

}

func ParseLine (line string,key string) (value string,err error) {

    var delimiters = []string {"="," ",":"}

    var quotes = []string {`\"`,`\'`}

    for dkey := range delimiters {

    }



    return "",nil

}

func SplitLine (line string ) (param string,value string ) {

    var delimiters = []string {"=",":"," "}
    var quotes = []string {`\"`,`\'`}
    var param_candidates  []string
    var value_candidates  []string

    for i := range delimiters {

        splitted_line := strings.Split(line,delimiters[i])

        if delimiters[i]!=" " {
            var stripped_line []string
            for word := range splitted_line {

                word=strings.Replace(word, " ", "", -1)

                stripped_line=append(stripped_line,word)

            }



        } else {





}
}
return "",""
}




