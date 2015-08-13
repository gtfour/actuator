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
    // DELL primary xml:
    // http://linux.dell.com/repo/community/content/el5-i386/repodata/primary.xml.gz

    test:=airparse.RepoFile {Url: "http://linux.dell.com/repo/community/content/el5-i386/repodata/primary.xml.gz"}

    test.Download()

//    test.Extract()

}


func main() {
    get_remote_file()
}


