package "agent"

type Agent struct {

    Id string
    Hostname string
    Ip string
    Architecture string
    Osname string
    Osfamily string
    Release string
    Configuration Configuration
    
}

type Repodirectory struct {

    Filepathes []Filepath
}

type Repofile struct {

    Filepath string

}

type Configuration struct {


}
