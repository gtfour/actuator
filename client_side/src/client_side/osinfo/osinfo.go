package main

var files []string = {"/etc/os-release":"rpm",
                      }



type Fedora struct {


}

type OpenSuSE struct {


}

type Ubuntu struct {


}

type OS struct {

    HostName string
    OsName string
    OsVersion string
    OsRelease string

}

func (os *OS)  CheckEtcFiles () err error {






}
