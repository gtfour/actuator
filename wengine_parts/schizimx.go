//
//
// check remote resource for any update 
//
//
package schizimx

import ("net/http";"io";"os")


func get_remote_file(){

    out,_:=os.Create("/tmp/Packages.gz")

    defer out.Close()

    resp,_:=http.Get("http://ubuntu-cloud.archive.canonical.com/ubuntu/dists/trusty-updates/kilo/main/binary-amd64/Packages.gz")

    defer resp.Body.Close()

    _, _=io.Copy(out, resp.Body)

}



