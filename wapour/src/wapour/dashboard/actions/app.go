package actions
import . "wapour/api/wengine"
import . "wapour/core/web/table"


func Actions() (string) {

    api := GetApi("","","")
    _,action_list:=api.ActionsList()
    table:=CreateTable()
    table.Name  = "Actions Table"
    table.Id    = "actions_table"
    table.Title = "Actions Table"
    for id := range  action_list {


        action:=action_list[id]
        row   :=Row{}
        row.Fields=[]string{ action.Name, action.Command}
        table.Rows= append( table.Rows, row )

    }
    return table.Render()


}
