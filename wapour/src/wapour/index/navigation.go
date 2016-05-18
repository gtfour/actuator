package index
import "github.com/gin-gonic/gin"

var JSON_URL_HEADER = "/dashboard/json"

type SubItem struct {
    Name     string
    Url      string
    Id       string
}

type Item struct {
    Name     string
    Icon     string
    SubItems []SubItem
    Id       string
}

func GetNavigationMenu()( []Item ){

    hostsList    := SubItem{Name:"list",Url:"hosts"}
    filesList    := SubItem{Name:"list",Url:"files"}
    filesAdd     := SubItem{Name:"add",Url:"#files/add"}
    actionsList := SubItem{Name:"list",Url:"actions"}
    actionsAdd  := SubItem{Name:"add",Url:"#actions/add"}
    wsMessages  := SubItem{Name:"messages",Url:"websocket"}

    return  []Item { Item { Name:"Hosts"   ,Icon: "fa-desktop",         SubItems:[]SubItem {hostsList}  },
                     Item { Name:"Files"   ,Icon: "fa-stack-overflow",  SubItems:[]SubItem {filesList,    filesAdd }  },
                     Item { Name:"Actions" ,Icon: "fa-gamepad",         SubItems:[]SubItem {actionsList, actionsAdd}  },
                      Item { Name:"WebSocket" ,Icon: "fa-flash",         SubItems:[]SubItem {wsMessages} }}
}

func GetNavigationMenuJson()( gin.H ){

    default_icon:="fa fa-circle-o"

    hostsList   :=gin.H{"name":"list","title":"List","url":JSON_URL_HEADER+"/hosts",  "icon":default_icon}
    filesList   :=gin.H{"name":"list","title":"List","url":JSON_URL_HEADER+"/files",  "icon":default_icon}
    actionsList :=gin.H{"name":"list","title":"List","url":JSON_URL_HEADER+"/actions","icon":default_icon}
    wsMessages  :=gin.H{"name":"messages","title":"Messages","url":JSON_URL_HEADER+"/websocket","icon":default_icon}

    hosts_item   := gin.H{"name":"Hosts",    "title":"Hosts"     ,"subitems":[]gin.H{hostsList}  , "icon":"fa-desktop"}
    files_item   := gin.H{"name":"Files",    "title":"Files"     , "subitems":[]gin.H{filesList}  , "icon":"fa-stack-overflow"}
    actions_item := gin.H{"name":"Actions",  "title":"Actions"   , "subitems":[]gin.H{actionsList}, "icon":"fa-gamepad"}
    ws_item      := gin.H{"name":"WebSocket","title":"WebSocket" , "subitems":[]gin.H{wsMessages}, "icon":"fa-flash"}

    menu_name:="Admin Dashboard"
    //navigation_menu["menu_name"]  = menu_name
    //navigation_menu["menu_items"] = []gin.H{hosts_item, files_item, actions_item, ws_item}

    return gin.H{"menu_name":menu_name, "menu_items":[]gin.H{hosts_item, files_item, actions_item, ws_item}}

}
