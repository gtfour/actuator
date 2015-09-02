package main

import (
        "bytes"
        "fmt"
        "log"
        "os/exec"
        "strings"
)


var packages_db_file = "/var/lib/rpm/Packages"
var environment_dir = "/var/lib/rpm"



func main() {
  packages,_:=PackagesList()
  fmt.Printf("%s",packages)
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
  packages= strings.Split(out.String(),"\n")
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
  cmd.Stdout = &out
  err = cmd.Run()
  if err != nil {
      log.Fatal(err)
  }
  lines= strings.Split(out.String(),"\n")


}


