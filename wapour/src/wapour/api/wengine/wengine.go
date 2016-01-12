package wengine

type Api struct {

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

func (a *Api) ActionsList()(err error, actions []Action) {

    actions = []Action {Action{Name:"partitions_list",Command:"cat /proc/partitions"}, Action{Name:"networking",Command:"ifconfig -a"}}

    return nil, actions


}

func (a *Api) FilesList()(err error, actions []Action) {

    actions = []Action {Action{Name:"partitions_list",Command:"cat /proc/partitions"}, Action{Name:"networking",Command:"ifconfig -a"}}

    return nil, actions


}
