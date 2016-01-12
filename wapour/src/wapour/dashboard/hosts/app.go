package hosts
import . "wapour/api/wengine"
import . "wapour/core/web/table"

func Hosts() (string) {

    api := GetApi("","","")
    _,hosts_list:=api.HostsList()
    table:=CreateTable()
    table.Name  = "Hosts Table"
    table.Id    = "hosts_table"
    table.Title = "Hosts Table"
    table.HeaderFields = []string {"ID"}
    for id := range hosts_list {


        host       := hosts_list[id]
        row        := Row{}
        row.Fields =  []string{ host.Id}
        table.Rows =  append( table.Rows, row )

    }
    return table.Render()


}

