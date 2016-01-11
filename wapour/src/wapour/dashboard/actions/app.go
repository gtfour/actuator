package actions
import . "wapour/api/wengine"
import . "wapour/core/web/table"


func Actions() (table Table) {

    api := GetApi("","","")
    _,actions:=api.ActionsList()
    table=CreateTable()
    for id := range  actions {


        action:=actions[id]
        row   :=Row{}
        row.Fields=[]string{ action.Name, action.Command}
        table.Rows= append( table.Rows, row )

    }
    return table


}
