//
//
// check remote resource for any update 
//
//
package main

//import ("net/http";"io";"os")
import "wengine_parts/airparse"
import "wengine_parts/dbsync"
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
    // http://127.0.0.1:8080/primary_dell.xml.gz

    test:=airparse.RepoFile {Url: "http://127.0.0.1:8080/primary_dell.xml.gz"}

    err:=test.Download()

    fmt.Println(err)

    dbsync.UploadStructToDb(test)



    if err==nil {

        for pkg_number:= range test.Packages{

            fmt.Println("-----")

            fmt.Printf("type: %s\n",test.Packages[pkg_number].Type)

            fmt.Printf("package_name: %s\n",test.Packages[pkg_number].Name)

            fmt.Printf("architecture: %s\n",test.Packages[pkg_number].Architecture)

            fmt.Printf("version: epoch: %s ver: %s rel: %s \n",test.Packages[pkg_number].PackageVersionField.Epoch,test.Packages[pkg_number].PackageVersionField.Ver,test.Packages[pkg_number].PackageVersionField.Rel)

            fmt.Printf("Checksum: %s Type: %s Pkgid: %s\n", test.Packages[pkg_number].PackageChecksumField.Checksum,test.Packages[pkg_number].PackageChecksumField.Type,test.Packages[pkg_number].PackageChecksumField.Pkgid)

            fmt.Printf("summary: %s\n",test.Packages[pkg_number].Summary)


            fmt.Printf("####\ndescription: %s\n#####\n",test.Packages[pkg_number].Description)

            fmt.Printf("packager: %s\n",test.Packages[pkg_number].Packager)

            fmt.Printf("package_url: %s\n",test.Packages[pkg_number].Url)

            fmt.Printf("time: file: %s build: %s\n",test.Packages[pkg_number].PackageTimeField.File,test.Packages[pkg_number].PackageTimeField.Build)

            fmt.Printf("size: package: %s installed: %s archive: %s\n",test.Packages[pkg_number].PackageSizeField.Package,test.Packages[pkg_number].PackageSizeField.Installed,test.Packages[pkg_number].PackageSizeField.Archive)

            fmt.Printf("location: %s\n",test.Packages[pkg_number].LocationHref.Href)

            fmt.Printf("location: %s\n",test.Packages[pkg_number].LocationHref.Href)

            fmt.Printf("  license : %s vendor: %s group: %s buildhost: %s sourcerpm: %s \n",test.Packages[pkg_number].RpmFormatField.License,test.Packages[pkg_number].RpmFormatField.Vendor,test.Packages[pkg_number].RpmFormatField.Group,test.Packages[pkg_number].RpmFormatField.Buildhost,test.Packages[pkg_number].RpmFormatField.SourceRpm)


            fmt.Println("::: relations :::")


            fmt.Println(":::           :::")

           

            fmt.Println("-----")


            }


    }
}


func main() {
    get_remote_file()
}


