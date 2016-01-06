package wengine

type Feature struct {


}

type Api struct {

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

    actions = []Action { Action{Name:"partitions_list",Command:"cat /proc/partitions"}}

    return


}
