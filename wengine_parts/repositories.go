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
    Installed-Size int
    Maintainer string
    Architecture string
    Version string
Depends: ceilometer-common (= 2015.1.0-0ubuntu1~cloud0), init-system-helpers (>= 1.13~), sysv-rc (>= 2.88dsf-24) | file-rc (>= 0.8.16)
, python:any
Supported: 18m
Filename: pool/main/c/ceilometer/ceilometer-agent-central_2015.1.0-0ubuntu1~cloud0_all.deb
Size: 4862
MD5sum: ada1d7f4bddcce586543de615f0d9eac
SHA1: 3b5b5faa0cdc2b2626938b941f9604cd33b53006
SHA256: df879c9da8ba89c593698f84c6a407e3d6a1cc091558eb5ff80e48020454c488
Description: ceilometer central agent
Description-md5: 69931ca0e99876c83677bf349c97e7bb



}


