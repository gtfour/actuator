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

    Type string
    Package string
    Architecture string
    Epoch string
    Version string
    Release string
    SHA256 string
    PackageID string
    Summary string
    Description string
    Packager string
    Url string
    TimeFile uint32
    TimeBuild uint32
    SizePackage uint64
    SizeInstalled uint64
    SizeArchive uint64
    LocationHref string
    License string
    Vendor string
    Group string
    Buildhost string
    SourceRpm string
    HeaderRangeStart uint64
    HeaderRangeEnd uint64
    RpmDependencyOrProvision RpmRelation

}

type RpmRelation struct {

    Name string
    FileName string
    Architecture string
    Flags string
    Epoch string
    Version string
    Release string

}

