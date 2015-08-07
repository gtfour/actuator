package airparse

import ("net/http";"errors";"compress/gzip";"fmt")
//import "io"

var err = errors.New("airparse side error")

type RepoFile struct {

    Id string
    FileName string
    Url string
    DataGZ *http.Response
    DataXML http.Response
    Type string // deb or rpm

}

func (repofile *RepoFile) Download(){

    repofile.DataGZ, err=http.Get(repofile.Url)
    
    defer repofile.DataGZ.Body.Close()
}

func (repofile *RepoFile) Extract() string{

    r, _ := gzip.NewReader(repofile.DataGZ.Body)

    fmt.Println(r)

}

func (repofile *RepoFile) parse(){

}

func (repofile *RepoFile) handle(){


}



