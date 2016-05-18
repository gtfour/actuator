package userspace
import "wapour/api/webclient"
import "github.com/gin-gonic/gin"

type SubItem struct {
    Name     string
    Url      string
    Id       string
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
        item      := Item{ Name:dgroup.Title, Id:dgroup.Id, Icon:dgroup.Icon}
        for dashboard_id := range dinstance.DashboardList.List {
            dashboard := dinstance.DashboardList.List[dashboard_id]
            if IsIn(dashboard.Id, dgroup.List){
                subitem:= SubItem{ Name:dashboard.Title, Url:"#"+dgroup.Id+"/"+dashboard.Url, Id:dashboard.Id}
                item.SubItems = append(item.SubItems, subitem)
            }
        }
        dgroup_list = append(dgroup_list, item)
    }
    return dgroup_list
}

func CreateNavigationMenuJson(dinstance webclient.DashboardListResult)( gin.H ){

    menu_name:="User Dashboard"

    //dgroup_list        := make([]Item,0)
    //default_item       := Item{Name:"Default",Id:"default_dashboard",Icon:"fa-cubes"}
    //var grouped_dashboards  []string
    var menu_items = []gin.H{}

    for dgroup_id := range dinstance.DashboardGroupList.List {
        dgroup    := dinstance.DashboardGroupList.List[dgroup_id]
        var item = gin.H{"name":dgroup.Name,"icon":dgroup.Icon,"title":dgroup.Title,"id":dgroup.Id, "subitems":[]gin.H{}}
        //dgroup    := dinstance.DashboardGroupList[dgroup_id]
        //item      := Item{ Name:dgroup.Title, Id:dgroup.Id, Icon:dgroup.Icon}
        var subitems = []gin.H{}
        for dashboard_id := range dinstance.DashboardList.List {

            dashboard := dinstance.DashboardList.List[dashboard_id]
            if IsIn(dashboard.Id, dgroup.List){
                //subitem:= SubItem{ Name:dashboard.Title, Url:"#"+dgroup.Id+"/"+dashboard.Url, Id:dashboard.Id}
                //item.SubItems = append(item.SubItems, subitem)
                var subitem = gin.H{"name":dashboard.Name, "id":dashboard.Id, "title":dashboard.Title, "icon":dashboard.Icon, "url":""}
                subitems = append(subitems, subitem)
            }
        }
        item["subitems"] = subitems
        //dgroup_list = append(dgroup_list, item)
        menu_items = append(menu_items, item)
    }
    //return dgroup_list
    return gin.H{"menu_name":menu_name, "menu_items":menu_items}
}


func IsIn(x string,xs []string)(bool) {
    for i := range xs {
        y:= xs[i]
        if y == x {return true}
    }
    return false
}
