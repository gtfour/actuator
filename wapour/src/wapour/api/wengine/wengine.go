package wengine

type Api struct {
    username   string
    password   string
    token      string
    auth_url   string
    auth_port  string
    auth_proto string

}

func GetApi(username string, password string, auth_host string) ( a Api ) {

    a = Api{}
    return a

}

func (a *Api) Login ()(err error) {
    return err
}

func (a *Api) Logout ()(err error) {
    return err
}

func (a *Api) HostsUp()(err error) {
    return err
}

func (a *Api) ActionsList()(err error, actions []_Action) {

    actions = []_Action {_Action{Name:"partitions_list",Command:"cat /proc/partitions"}, _Action{Name:"networking",Command:"ifconfig -a"}}

    return nil, actions


}

func (a *Api) FilesList()(err error, files []_File) {

    files = []_File {_File{Name:"mountpoint_list",Path:"/etc/fstab",IsDir:false}, _File{Name:"logs",Path:"/var/log",IsDir:true}}

    return nil, files


}

func (a *Api) HostsList()(err error, hosts []_Host) {


    hosts =  []_Host {_Host{Id:"f977dcf2"}, _Host{Id:"k7qzyxlq"}, _Host{Id:"0vswd9io"}, _Host{Id:"fzdqmim6"}}

    return nil, hosts

}
