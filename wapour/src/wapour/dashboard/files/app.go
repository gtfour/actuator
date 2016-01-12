package files
import . "wapour/api/wengine"
import . "wapour/core/web/table"


func Files() (string) {

    api := GetApi("","","")
    _,files_list:=api.FilesList()
    table:=CreateTable()
    table.Name  = "Files Table"
    table.Id    = "files_table"
    table.Title = "Files Table"
    for id := range files_list {


        action:=action_list[id]
        row   :=Row{}
        row.Fields=[]string{ action.Name, action.Command}
        table.Rows= append( table.Rows, row )

    }
    return table.Render()


}
