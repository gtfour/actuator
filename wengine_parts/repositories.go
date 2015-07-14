package repository

type Repository struct {
    Id         int
    Name int
    Path     int64 // 
    Type     string // deb or rpm
    Markers  []string
    Packages []Package
}

type Package struct {
    Id         int
    Name int
    Path     int64 //
    Type     string // deb or rpm
    Markers  []string
}



type DebPackage struct {

    Package string
    Source string
    Priority string
    Section string
    InstalledSize int
    Maintainer string
    Architecture string
    Version string
    Depends []DependencyDeb
    Supported string
    Filename string 
    Size int
    MD5sum string
    SHA1 string 
    SHA256 string 
    Description string 
    Descriptionmd5 string 
}

type DependencyDeb struct {

    Package string
    Version string
    Substitutes []string  

}



type RpmPackage struct {




}

