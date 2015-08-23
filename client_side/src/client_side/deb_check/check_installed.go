package deb_check

type StatusEntry struct {

    Name string
    Version string
    Architecture string
    Complete bool

}

type StatusFile struct {

    FilePath string
    Sum string
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

// main info /var/lib/dpkg/status
