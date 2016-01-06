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

    filesList    := SubItem{Name:"list",Url:"#files/list"}
    filesAdd     := SubItem{Name:"add",Url:"#files/add"}
    actionsList := SubItem{Name:"list",Url:"#actions/list"}
    actionsAdd  := SubItem{Name:"add",Url:"#actions/add"}

    return  []Item { Item { Name:"Files"    ,Icon: "fa-stack-overflow" ,SubItems:[]SubItem {filesList,    filesAdd }  },
                     Item { Name:"Actions" ,Icon: "fa-gamepad",         SubItems:[]SubItem {actionsList, actionsAdd}  }}

}
