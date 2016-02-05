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

    dgroup_list:= make([]Item,0)

    for dgroup_id := range dinstance.DashboardGroupList.List {
        dgroup    := dinstance.DashboardGroupList.List[dgroup_id]
        //dgroup    := dinstance.DashboardGroupList[dgroup_id]
        item      := Item{Name:dgroup.Title, Id:dgroup.Id, Icon:dgroup.Icon}
        for dashboard_id := range dinstance.DashboardList.List {
            dashboard := dinstance.DashboardList.List[dashboard_id]
            subitem:= SubItem{Name:dashboard.Name}
            item.SubItems = append(item.SubItems, subitem)
        }
        dgroup_list = append(dgroup_list, item)
    }
    return dgroup_list
}


