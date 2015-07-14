package repository

type Repository struct {
    Id         int
    Name int
    Path     int64 // 
    Type     string // deb or rpm
    Markers  []string
}

type Package struct {
    Id         int
    Name int
    Path     int64 //
    Type     string // deb or rpm
    Markers  []string
}


