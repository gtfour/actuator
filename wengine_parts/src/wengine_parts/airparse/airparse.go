package airparse

import ("net/http";"errors";"compress/gzip";"fmt")
//import "bytes"
import "reflect"
//import "xml"
import "io/ioutil"


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

    fmt.Println("Download block")

    fmt.Println(reflect.TypeOf(repofile.DataGZ.Body))

    fmt.Println("==============")

    reader, _ :=gzip.NewReader(repofile.DataGZ.Body)

    xml.Unmarshal(reader, &q)

    text, _:=ioutil.ReadAll(reader)

    

    fmt.Println("---")

    fmt.Println("%x",text)

    fmt.Println("---")

    defer repofile.DataGZ.Body.Close()
}

func (repofile *RepoFile) Extract(){

    r, _ := gzip.NewReader(repofile.DataGZ.Body)

    fmt.Println("---")
    fmt.Println(r)
    fmt.Println("---")

}

func (repofile *RepoFile) Parse(){

}

func (repofile *RepoFile) Handle(){


}



