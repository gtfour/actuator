package dashboard

type Dashboard struct {

    datasource DataSource

}

type Table struct {


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


func CreateDataSource (dstype string , )

