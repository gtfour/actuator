package cuda
import "client/cross"
import "client/common/types"

var database = cross.Database

func Preparing()(err error){
    err:=CheckDynimaTable()
    if err!=nil {
        err:=CreateDynimaTable()
        if err!=nil {
            return err
        }
    }else{
        return nil
    }
}

func CheckDynimaTable()(err error){
    check_query:=cross.Query{Type:types.CHECK_TABLE_EXIST, Table:types.DYNIMAS_T}
    _,err:=database.RunQuery(check_query)
    return err
}

func CreateDynimaTable()(err error){
    create_query:=cross.Query{Type:types.CREATE_NEW_TABLE}
    _,err:=database.RunQuery(create_query)
    return err
}
