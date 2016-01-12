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

    hostsList    := SubItem{Name:"list",Url:"hosts"}
    filesList    := SubItem{Name:"list",Url:"files"}
    filesAdd     := SubItem{Name:"add",Url:"#files/add"}
    actionsList := SubItem{Name:"list",Url:"actions"}
    actionsAdd  := SubItem{Name:"add",Url:"#actions/add"}

    return  []Item { Item { Name:"Hosts"   ,Icon: "fa-desktop",         SubItems:[]SubItem {hostsList}  },
                     Item { Name:"Files"   ,Icon: "fa-stack-overflow",  SubItems:[]SubItem {filesList,    filesAdd }  },
                     Item { Name:"Actions" ,Icon: "fa-gamepad",         SubItems:[]SubItem {actionsList, actionsAdd}  }}
}
