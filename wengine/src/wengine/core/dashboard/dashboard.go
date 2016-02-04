package dashboard

type DashboardList struct {

    List []Dashboard

}

type DashboardGroup struct{
    Title string
    Icon  string
    List  []string
}

type Dashboard struct {
    Id           string
    Name         string
    Title        string
    Composition  []Component
}

type ComponentSet struct {


}

type Component interface {
    GetType(string)
    GetData([][]string)

}

type Table struct {
    datasource DataSource
}

type Action struct {


}

type File struct {


}

type DataSource interface {
    GetData(Data)
}

type Data interface {


}

func NewDashboard(name string)(d Dashboard) {
    return d
}

func ( d *Dashboard ) GetData ()(xdata interface {}) {

    return
}


func CreateDataSource (dstype string ) {


}

