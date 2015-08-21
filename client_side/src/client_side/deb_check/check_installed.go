package deb_check

type DebPackageStatusEntry struct {

    Name string
    Version string
    Architecture string

}

type DpkgStatusFile struct {

    FilePath string
    Sum string

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
