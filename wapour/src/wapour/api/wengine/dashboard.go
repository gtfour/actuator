package wengine

var TABLE_TYPE_COMPONENT     int = 1000
var CHART_TYPE_COMPONENT     int = 1001

var TABLE_DATA_CREATE_ACTION int = 2000
var TABLE_DATA_EDIT_ACTION   int = 2001
var TABLE_DATA_DELETE_ACTION int = 2002
var TABLE_DATA_LINK_ACTION   int = 2003

type DashboardList      struct{
    List []Dashboard
}
type DashboardGroupList struct{
    List []DashboardGroup
}



type Dashboard struct {

    Id           string
    Name         string
    Title        string
    Url          string
    Composition  []Component

}

type DashboardGroup struct{
    Id    string
    Title string
    Icon  string
    List  []string
}




type ComponentSet struct {


}

type Component interface {
    GetType(string)
    GetData([][]string)

}



type _Table struct {

    Name         string
    TableActions []_TableAction
    RowActions   []_RowAction



}

type _TableAction struct {

}
type _RowAction struct {

}


func (a *Api) DashboardList()(err error, dashboards []Dashboard) {

    dashboards = []Dashboard {Dashboard{Name:"mountpoints",Title:"Mountpoints",Id:"ux4bxa2nscr3bsmm"}, Dashboard{Name:"network_settings",Title:"Network Settings",Id:"pdjku29gr9x2naq8"}}

    return nil, dashboards


}

func (d *Dashboard) GetData () {



}


