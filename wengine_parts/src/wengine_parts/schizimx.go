//
//
// check remote resource for any update 
//
//
package main

//import ("net/http";"io";"os")
import "wengine_parts/airparse"


func get_remote_file() {

    //out,_:=os.Create("/tmp/Packages.gz")
    //defer out.Close()
    //resp,_:=http.Get("http://ubuntu-cloud.archive.canonical.com/ubuntu/dists/trusty-updates/kilo/main/binary-amd64/Packages.gz")
    //defer resp.Body.Close()
    //_, _=io.Copy(out, resp.Body)

    test:=airparse.RepoFile {Url: "http://download.opensuse.org/distribution/13.2/repo/oss/suse/repodata/appdata.xml.gz"}

    test.Download()

//    test.Extract()

}


func main() {
    get_remote_file()
}


