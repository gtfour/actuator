package main

import (
        "bytes"
        "fmt"
        "log"
        "os/exec"
        "strings"
)




func main() {
  packages,_:=PackagesList()
  fmt.Printf("%s",packages)
  for pkg_num:=range packages { info,_:=GetInfo(packages[pkg_num]) ; fmt.Printf("%s %s %s %s\n",info.Name,info.Version,info.Release,info.Architecture)}
}

func PackagesList () (packages []string,err error){

  cmd := exec.Command("rpm", "-qa","--queryformat='%{NAME}\n'")
  cmd.Stdin = strings.NewReader("some input")
  var out bytes.Buffer

  cmd.Stdout = &out
  err = cmd.Run()
  if err != nil {
      log.Fatal(err)
  }
  packages_temp:= strings.Split(out.String(),"\n")
  for i:=range packages_temp { pkg_name:=packages_temp[i] ;  pkg_name=strings.Replace(pkg_name, `'`, "", -1) ; pkg_name=strings.Replace(pkg_name, " ", "", -1) ; packages=append(packages,pkg_name) }
  return packages,err


}

type Info struct {
  Name string
  Version string
  Release string
  Architecture string
}

func GetInfo(package_name string) (info Info,err error) {

  cmd := exec.Command("rpm", "-qi",package_name)
  cmd.Stdin = strings.NewReader("some input")
  var out bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &out
  cmd.Stderr = &stderr
  err = cmd.Run()
  if err != nil {
      log.Fatal(err)
  }
  fmt.Printf("\nErr: %q\n",stderr.String())
  lines:=strings.Split(out.String(),"\n")
  for line_num :=range lines {
    line:=lines[line_num]
    if strings.HasPrefix(line, "Name") { sp_st:= strings.Split(line,":");  if (len(sp_st)==2) {info.Name=strings.Replace(sp_st[1], " ", "", -1) }} 
    if strings.HasPrefix(line, "Version") { sp_st:= strings.Split(line,":");  if (len(sp_st)==2) {info.Version=strings.Replace(sp_st[1], " ", "", -1) }}
    if strings.HasPrefix(line, "Release") { sp_st:= strings.Split(line,":");  if (len(sp_st)==2) {info.Release=strings.Replace(sp_st[1], " ", "", -1) }}
    if strings.HasPrefix(line, "Architecture") { sp_st:= strings.Split(line,":");  if (len(sp_st)==2) {info.Architecture=strings.Replace(sp_st[1], " ", "", -1) }}
  }
  return
  // strings.HasPrefix(line, "Version"))
}



