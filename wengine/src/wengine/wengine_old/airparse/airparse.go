package airparse

import ("net/http";"errors";"compress/gzip")
//import "bytes"
//import "reflect"
import "encoding/xml"
import "io/ioutil"
import "wengine/repository"
// RpmMetadata 


var err = errors.New("airparse side error")

type RepoFile struct {

    Id       string
    FileName string
    Url      string
    DataGZ   *http.Response
    DataXML  http.Response
    Type     string // deb or rpm
    Packages []repository.RpmPackage

}

func (repofile *RepoFile) Download() (err error){

    repofile.DataGZ, err=http.Get(repofile.Url)

    if err!=nil {

         return err

    }

    reader, err :=gzip.NewReader(repofile.DataGZ.Body)

    if err!=nil {

        return err

    }

    text, err:=ioutil.ReadAll(reader)

    if err!=nil {

        return err

    }

    repometadata_struct := repository.RpmMetadata {}

    err=xml.Unmarshal(text, &repometadata_struct)

    if err!=nil {

        return err
    }

    repofile.Packages=repometadata_struct.RpmPackages

    defer repofile.DataGZ.Body.Close()

    return nil
}

func (repofile *RepoFile) Extract(){

    //r, _ := gzip.NewReader(repofile.DataGZ.Body)

}

func (repofile *RepoFile) Parse(){

}

func (repofile *RepoFile) Handle(){


}



