package repository

import "encoding/xml"

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

    Package string  `xml:package`
    Type string `xml:"type"`
    Name string `xml:"name"`
    Architecture string `xml:"arch"`
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

type RpmMetadata struct {

     Metadata xml.Name `xml:"Metadata"`
     XmlNS string `xml:"xmlns,attr"`
     XmlNsSuSE string `xml:"xmlns:suse,attr"`
     XmlNsRpm string `xml:"xmlns:rpm,attr"`
     PackagesCount string `xml:"packages,attr"`
     RpmPackages []RpmPackage

}

type RpmMetadataField struct{


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

