package main
// Determines current operating system

import "fmt"
import "io/ioutil"
import "strings"
// I am glad to introduce you "f-shit power" and "How does this fking code work" technologies 
// I have to to finish handling of vp array !!!



type OS struct {

    Hostname string
    Name string
    Version string
    Release string
    Files []string
    VirtualProvider []string // по задумке сюда будут складываться строки начиная со второй, которые были найдены
                             // проходом значений хеша по ключу директ 
                             // по идее директ должен включать файлы имеющие всего одну строку, не содержащую ключей а только одно значение
                             // прошу прощения за этот бред )

}



func main() {

   operating_system:=&OS{}


   operating_system.GetHostname()
   operating_system.GetName()
   operating_system.GetVersion()
   operating_system.GetRelease()


   fmt.Printf("hostname %s",operating_system.Hostname)
   fmt.Printf("name %s",operating_system.Name)
   fmt.Printf("version %s",operating_system.Version)
   fmt.Printf("release %s",operating_system.Release)
   fmt.Printf("------------------------------")
   fmt.Println(operating_system.VirtualProvider)

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

    // proccessing lines from virtual provider 
    var empty_vp []string
    for line_num :=range os.VirtualProvider{
        SingleLineProcessing(&values,&empty_vp,os.VirtualProvider[line_num],complex_keys,false)
    }
    os.Hostname,_=ValidateValue(values,key)

    return nil

}

func (os *OS) GetName() (err error) {


    var providers = map[string][]string {"complex":{"/etc/SuSE-release","/etc/SuSE-brand" , "/etc/lsb-release", "/etc/os-release"},"direct":{"/etc/redhat-release", "/etc/fedora-release","/etc/SuSE-brand"}}

    var complex_keys = []string {"NAME","DISTRIB_ID"}

    key:="name"
    var values []string

    values,_,os.VirtualProvider=GetParamValue(providers,complex_keys)
    // proccessing lines from virtual provider
    var empty_vp []string
    for line_num :=range os.VirtualProvider{
        SingleLineProcessing(&values,&empty_vp,os.VirtualProvider[line_num],complex_keys,false)
    }

    os.Name,_=ValidateValue(values,key)


    return nil



}

func (os *OS) GetVersion() (err error) {

    var providers = map[string][]string {"complex":{"/etc/SuSE-release", "/etc/lsb-release", "/etc/os-release", "/etc/SuSE-brand","/etc/issue","/etc/redhat-release"}}

    var complex_keys = []string {"VERSION_ID","VERSION","release","DISTRIB_RELEASE","ID"}

    key:="version"
    var values []string
    values,_,os.VirtualProvider=GetParamValue(providers,complex_keys)
    // proccessing lines from virtual provider
    var empty_vp []string
    for line_num :=range os.VirtualProvider{
        SingleLineProcessing(&values,&empty_vp,os.VirtualProvider[line_num],complex_keys,false)
    }

    os.Version,_=ValidateValue(values,key)


    return nil

}

//

func (os *OS) GetRelease() (err error) {

    var providers = map[string][]string {"complex":{"/etc/lsb-release","/etc/fedora-release","/etc/redhat-release","/etc/SuSE-brand","/etc/os-release","/etc/issue"}}

    var complex_keys = []string {"release","DISTRIB_RELEASE","VERSION_ID","VERSION","VERSION_ID","RELEASE"}

    key:="release"
    var values []string
    values,_,os.VirtualProvider=GetParamValue(providers,complex_keys)
    // proccessing lines from virtual provider
    var empty_vp []string
    for line_num :=range os.VirtualProvider{
        SingleLineProcessing(&values,&empty_vp,os.VirtualProvider[line_num],complex_keys,false)
    }

    os.Release,_=ValidateValue(values,key)



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

                lines,_:=ReadFileLines(filename)

                    for num := range lines {

                        //value,vp_new,_:=ParseLine(lines[num],complex_keys)
                        //if len(vp_new)>0 { for i:= range vp_new { vp=append(vp,vp_new[i]) } }
                        //possible_value_candidates=append(possible_value_candidates,value)
                        var vp_new []string
                        SingleLineProcessing(&possible_value_candidates,&vp_new,lines[num],complex_keys,true)
                        if len(vp_new)>0 { for i:= range vp_new { vp=append(vp,vp_new[i]) } }


                    }


            }
       }


    }
    return possible_value_candidates,nil,vp


}


func SingleLineProcessing (possible_value_candidates *[]string,vp *[]string,line string,complex_keys []string,collect_vp bool){

    // 
    // if collect_vp == false thats means  we are working with VP
    // 
    var value  string 
    value,vp_new,_:=ParseLine(line , complex_keys)
    if (len(vp_new)>0 && collect_vp==true)  { for i:= range vp_new { *vp=append(*vp,vp_new[i]) } }
    *possible_value_candidates=append(*possible_value_candidates,value)
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

    if key == "hostname" { fmt.Printf("hostname:")  ;  for i := range values { fmt.Printf("%s|",values[i]) } ; fmt.Println("\n")  }
    if key == "name" { fmt.Printf("name:")  ; for i := range values { fmt.Printf("%s|",values[i]) } ; fmt.Println("\n")  }
    if key == "version" { fmt.Printf("version:")  ; for i := range values { fmt.Printf("%s|",values[i]) }  ; fmt.Println("\n") }
    if key == "release" { fmt.Printf("release:")  ; for i := range values { fmt.Printf("%s|",values[i]) } ; fmt.Println("\n")  }

    for i := range values{

       if key == "hostname" {


           if (len(values[i])>len(value))&&(! strings.HasPrefix(values[i], "local")) { value=values[i] }

       }
       if key == "name" {

           //var exp_values []string
           for i := range values {

               if !IsContainDigit(values[i]) { }

           }

    }
    }
    return value, nil


}
func ParseLine (line string,complex_keys []string) (value string,vp []string, err error) {

    var param string


    param,value = SplitLine(line)
    for key := range complex_keys { if strings.EqualFold(complex_keys[key],param) {  fmt.Printf("Debug::  cckey: %s param:  %s\n",complex_keys[key],param )  ; return value,vp,nil } }
    // below extension to parse redhat-release and SuSE-brand txt files




    if (param==value) {
                      name:="NAME="
                      version:="VERSION="
                      release:="RELEASE="
                      var release_word_number int
                      var name_len int
                      sp_line:= strings.Split(value," ")
                      for wid := range sp_line {

                          if sp_line[wid] == "release" {
                               release_word_number=wid
                               name_len=wid-1
                               if len(sp_line[:name_len])>0 {name=name+strings.Join(sp_line[:name_len]," ")}
                               if len(sp_line)>(release_word_number+1) {
                                   version=version+sp_line[release_word_number+1] ;
                                   release=release+"1"
                               }

                          }
                          if strings.Index(sp_line[wid], ".")>=0 {
                              name_len=wid-1 ;
                              sp_line[wid]=strings.Replace(sp_line[wid], "(", "", -1)
                              sp_line[wid]=strings.Replace(sp_line[wid], ")", "", -1)
                              version_and_release := strings.Split(sp_line[wid],".")
                              fmt.Printf("\nVERSION_AND_RELEASE %s\n",version_and_release)
                              if (len(version_and_release)>=2) {
                                  version=version+string(version_and_release[0])
                                  // test
                                  //release=release+string(version_and_release[1])
                                  if (len(version_and_release[1:])>1) {

                                      release=release+string(strings.Join(version_and_release[1:],".")) ; fmt.Printf("\nif: %s\b",release)  } else {

                                      release=release+string(version_and_release[1]) }
                                      fmt.Printf("\nelse: %s\n",release)
                              } }

                       }

    if name!="NAME=" {vp=append(vp,name)}
    if version!="VERSION=" {vp=append(vp,version)}
    if release!="RELEASE=" {vp=append(vp,release)}
    }
    return "",vp,nil

}

func SplitLine (line string ) (param string,value string ) {

    var delimiters = []string {":","="," "}

    for i := range delimiters {

        splitted_line := strings.Split(line,delimiters[i])

        if delimiters[i]!=" " {

            var stripped_line []string
            var subwords_splitted_by_space []string
            var subwords_line string

            if len(splitted_line) > 1 {
            for word_num := range splitted_line {

                word:=splitted_line[word_num]

                word=strings.Replace(word, `"`, "", -1) // -1 means that Replace should replace all space entries

                word=strings.Replace(word, `'`, "", -1) // if define 2 as last arg it will replace two times

                stripped_line=append(stripped_line,word)



                subwords_splitted_by_space:=strings.Split(word," ")

                subwords_line=strings.Join(subwords_splitted_by_space," ")
                //fmt.Printf("stripped_line: %s subwords_line:  %s \n",stripped_line,subwords_line)



            } }
            if len(stripped_line)==2 && len(subwords_splitted_by_space)<=1 { param = stripped_line[0] ; value = stripped_line[1]  }

            if len(stripped_line)==2 && len(subwords_splitted_by_space)>1  { param = stripped_line[0] ; value = subwords_line  }



        } else {  if (param=="" && value == "") { param=line ; value=line } }
    }
    fmt.Printf("Input: %s Param: %s Value: %s\n",line,param,value)
    return param,value
}



func IsContainDigit(line string)(b bool){ 

    for char:=range line { if line[char] >= '0' && line[char] <= '9' {  b=true } } 
    return b

}
