package repository

type Repository struct {

    Id string
    Url string //"http://linux.dell.com/repo/community/content/el5-i386/"
    Type string // deb or rpm
    Mtime string

}

type RpmRepository struct {

    Repository
    PrimaryXml string "repodata/primary.xml.gz"

}

type DebRepository struct {

    Repository
    Packages string "repodata/Packages.gz"

}




type Info struct {


}

type RepoMarker struct {



}

