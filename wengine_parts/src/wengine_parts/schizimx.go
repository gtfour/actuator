//
//
// check remote resource for any update 
//
//
package main

//import ("net/http";"io";"os")
import "wengine_parts/airparse"
import "fmt"


func get_remote_file() {

    //out,_:=os.Create("/tmp/Packages.gz")
    //defer out.Close()
    //resp,_:=http.Get("http://ubuntu-cloud.archive.canonical.com/ubuntu/dists/trusty-updates/kilo/main/binary-amd64/Packages.gz")
    //defer resp.Body.Close()
    //_, _=io.Copy(out, resp.Body)
    // DELL primary xml:
    // http://linux.dell.com/repo/community/content/el5-i386/repodata/primary.xml.gz
    // local DELL repo xml:
    // http://127.0.0.1:8080/

    test:=airparse.RepoFile {Url: "http://linux.dell.com/repo/community/content/el5-i386/repodata/primary.xml.gz"}

    test.Download()

    for pkg_number:= range test.Packages{

        fmt.Println("===============================")
        fmt.Printf("type:%s\n",test.Packages[pkg_number].Type)
        fmt.Printf("package_name:%s\n",test.Packages[pkg_number].Name)
        fmt.Printf("architecture:%s\n",test.Packages[pkg_number].Architecture)
        fmt.Printf("version: epoch:%s ver:%s rel:%s \n",test.Packages[pkg_number].PackageVersionField.Epoch,test.Packages[pkg_number].PackageVersionField.Ver,test.Packages[pkg_number].PackageVersionField.Rel)
        fmt.Printf("Checksum: %s Type:%s Pkgid:%s ", test.Packages[pkg_number].PackageChecksumField.Checksum,test.Packages[pkg_number].PackageChecksumField.Type,test.Packages[pkg_number].PackageChecksumField.Pkgid)
        fmt.Printf("package_url:%s\n",test.Packages[pkg_number].Url)
        fmt.Printf("location:%s\n",test.Packages[pkg_number].LocationHref.Href)
        
        fmt.Println("===============================")


    }
}


func main() {
    get_remote_file()
}


