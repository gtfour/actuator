package index

type SubItem struct {
    Name     string
    Url      string
}

type Item struct {
    Name     string
    Icon     string
    SubItems []SubItem
}

func GetNavigationMenu()( []Item ){

    logsList    := SubItem{Name:"list",Url:"#logs/list"}
    logsAdd     := SubItem{Name:"add",Url:"#logs/add"}
    actionsList := SubItem{Name:"list",Url:"#actions/list"}
    actionsAdd  := SubItem{Name:"add",Url:"#actions/add"}

    return  []Item { Item { Name:"Logs"    ,Icon: "fa-stack-overflow" ,SubItems:[]SubItem {logsList,    logsAdd   }  },
                     Item { Name:"Actions" ,Icon: "fa-gamepad",SubItems:[]SubItem {actionsList, actionsAdd}  }}

}
