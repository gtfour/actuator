package files

import . "wapour/api/wengine"
import . "wapour/core/web/table"
import "fmt"


func Files() (string) {

    api := GetApi("","","")
    _,file_list:=api.FilesList()
    table:=CreateTable()
    table.Name  = "Files Table"
    table.Id    = "files_table"
    table.Title = "Files Table"
    table.HeaderFields = []string {"Name","Path","IsDir"}
    for id := range  file_list {


        file:=file_list[id]
        row   :=Row{}
        row.Fields=[]string{ file.Name, file.Path, fmt.Sprintf("%v",file.IsDir)}
        table.Rows= append( table.Rows, row )

    }
    return table.Render()


}

