package userspace
import "wapour/api/webclient"

type SubItem struct {
    Name     string
    Url      string
}

type Item struct {
    Name     string
    Id       string
    Icon     string
    SubItems []SubItem
}

func CreateNavigationMenu(dinstance webclient.DashboardListResult)( []Item ){

    dgroup_list        := make([]Item,0)
    //default_item       := Item{Name:"Default",Id:"default_dashboard",Icon:"fa-cubes"}
    //var grouped_dashboards  []string

    for dgroup_id := range dinstance.DashboardGroupList.List {
        dgroup    := dinstance.DashboardGroupList.List[dgroup_id]
        //dgroup    := dinstance.DashboardGroupList[dgroup_id]
        item      := Item{ Name:dgroup.Title, Id:dgroup.Id, Icon:dgroup.Icon }
        for dashboard_id := range dinstance.DashboardList.List {
            dashboard := dinstance.DashboardList.List[dashboard_id]
            if IsIn(dashboard.Id, dgroup.List){
                subitem:= SubItem{Name:dashboard.Title}
                item.SubItems = append(item.SubItems, subitem)
            }
        }
        dgroup_list = append(dgroup_list, item)
    }
    return dgroup_list
}

func IsIn(x string,xs []string)(bool) {
    for i := range xs {
        y:= xs[i]
        if y == x {return true}
    }
    return false
}
