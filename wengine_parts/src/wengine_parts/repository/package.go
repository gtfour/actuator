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
    PackageVersionField 
    PackageChecksumField
    Summary string `xml:"summary"`
    Description string `xml:"description"`
    Packager string `xml:"packager"`
    Url string `xml:"url"`
    PackageTimeField
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

type PackageVersionField struct{

    Epoch string `xml:"epoch,attr"`
    Ver string `xml:"ver,attr"`
    Rel string `xml:"rel,attr"`

}

type PackageChecksumField struct{


    Checksum string `xml:",chardata"`
    Type string `xml:"type,attr"`
    Pkgid string `xml:"pkgid,attr"`

}

type PackageTimeField struct{

    File string `xml:"file,attr"`
    Build string `xml:"build,attr"`

}

type PackageSizeField struct {

    Package string `xml:"package,attr"`
    Installed string `xml:"installed,attr"`
    Archive string `xml:"archive,attr"`



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

