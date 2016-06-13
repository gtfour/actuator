package cross

import "fmt"
import "errors"

var dynima_edit_error =  errors.New("Unable to edit dynima")

type Dynima struct {
    //parsers
    Id              string
    ids             []int // id column number
    SourcePath      string
    SourceType      string
    header          []string
    data            [][]string
    filters         []string //FilterList
    template        string
    data_indexes    [][]int
    delim_indexes   [][]int
}


func (d *Dynima) BindFilter (filter_name string)(error) {
    return nil
}

func (d *Dynima) UnbindFilter (filter_name string)(error) {
    return nil
}

//func (d *Dynima) RunFilter (filter_name string)(error) {
//    return nil
//}

func (d *Dynima) Save ()(error) {
    return nil
}

func (d *Dynima) GetData ()(error) {
    return nil
}

func (d *Dynima) SetTemplate()(error) {
    return nil
}

func (d *Dynima) SetSource (sourceType string, sourcePath string)(error) {
    return nil
}


func CreateDynima(id string)(error) {

    fmt.Printf("Storage error %v\nNew dynima id %s\n",STORAGE_INSTANCE.Error,id)
    return nil

}

func EditDynima(d Dynima)(error){
    if STORAGE_INSTANCE.Error == false {

        db:=STORAGE_INSTANCE.Db
        err=db.Update(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(db.dynimasTableName))
            if b==nil{ return nil }
            session,err:=b.CreateBucket([]byte(s.SessionId))
            if err==nil || err==bolt.ErrBucketExists { // If the key exist then its previous value will be overwritten
                err=session.Put([]byte("dashboard_id"),[]byte(s.DashboardId))
                if err!=nil{ return err }
                err=session.Put([]byte("user_id"),[]byte(s.UserId))
                if err!=nil{ return err }
                err=session.Put([]byte("token_id"),[]byte(s.TokenId))
                if err!=nil{ return err }
            } else { return err }
            return nil
        });

    }
    return dynima_edit_error
}
