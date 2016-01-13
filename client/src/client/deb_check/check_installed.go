package deb_check
import "runtime"

type StatusEntry struct {

    Name         string
    Version      string
    Architecture string
    Complete     bool
    Installed    bool

}

type StatusFile struct {

    FilePath          string
    Sum               string
    InstalledPackages []StatusEntry

}

func OpenStatusFile() {

}

func Watch() {


}

func CheckSum() {


}

func GenSumModifiedEvent() {


}

func SendJsonData() {



}

func GetOsProp ()  {



}







// main info /var/lib/dpkg/status
