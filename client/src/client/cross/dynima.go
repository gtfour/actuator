package cross

import "fmt"
import "encoding/json"
import "github.com/boltdb/bolt"


type Dynima struct {
    //parsers
    Id              string
    ids             []int // id column number
    SourcePath      string
    SourceType      string // file/dir/command
    header          []string
    Data            [][]string
    filters         []string //FilterList
    template        string // will automatically generated based on delim_indexes
    lockTemplate    bool
    data_indexes    [][]int
    delim_indexes   [][]int
}


func (d *Dynima) BindFilter (filter_name string)(error) {
    return nil
}

func (d *Dynima) UnbindFilter (filter_name string)(error) {
    return nil
}


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

func (d *Dynima) Write()(err error){
    if STORAGE_INSTANCE.Error == false {
        db:=STORAGE_INSTANCE.Db
        err=db.Update(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(STORAGE_INSTANCE.dynimasTableName))
            if b==nil{ return collection_open_error }
            encoded, err := json.Marshal(d)
            if err!=nil{ return err }
            return b.Put([]byte(d.Id), encoded) //CreateBucket has been replaced to CreateBucketIfNotExists because when err==bolt.ErrBucketExists - dynima is nil
        });

    }
    return collection_entry_edit_error
}

func (d *Dynima) UpdateTemplate()(){
    if d.delim_indexes != nil {
    }
}

func EditDynimaData(d Dynima, data [][]string)(err error){
    if STORAGE_INSTANCE.Error == false {

        db:=STORAGE_INSTANCE.Db
        err=db.Update(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(STORAGE_INSTANCE.dynimasTableName))
            if b==nil{ return collection_open_error }
            dynima,err:=b.CreateBucketIfNotExists([]byte(d.Id)) //CreateBucket has been replaced to CreateBucketIfNotExists because when err==bolt.ErrBucketExists - dynima is nil
            if err==nil || err==bolt.ErrBucketExists { // If the key exist then its previous value will be overwritten
                //fmt.Printf("\nEdit dynima:\n%v\nError: %v\n",d,err)
                err=dynima.Put([]byte("source_path"),[]byte(d.SourcePath))
                if err!=nil{ return err }
                //
            } else { return err }
            return nil
        });

    }
    return collection_entry_edit_error
}

func RemoveDynima(id string)(error){
    if STORAGE_INSTANCE.Error == false {
        db:=STORAGE_INSTANCE.Db
        err := db.Update(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(STORAGE_INSTANCE.dynimasTableName))
            if b==nil{ return collection_open_error }
            err:=b.Delete([]byte(id))
            return err
        });
        return err
    }
    return collection_entry_remove_error
}


func GetDynima(id string)(*Dynima,error){
    var err error
    if STORAGE_INSTANCE.Error == false {
        db     := STORAGE_INSTANCE.Db
        dynima := &Dynima{}
        err = db.View(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(STORAGE_INSTANCE.dynimasTableName))
            if b==nil{ return collection_open_error }
            data:=b.Get([]byte(id))
            err = json.Unmarshal(data, &dynima)
            return err
        });
        if err == nil { return dynima, nil } else { return nil, collection_entry_get_error }
    }
    return nil, collection_entry_get_error
}

func GetDynimasByPath(path string) ( dynimas  []Dynima , err  error ) {

    if STORAGE_INSTANCE.Error == false {
        db     := STORAGE_INSTANCE.Db
        dynimas = make([]Dynima,0)
        err = db.View(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(STORAGE_INSTANCE.dynimasTableName))
            if b==nil{ return collection_open_error }
            err=b.ForEach(func(key, value []byte)(error){
                dynima := Dynima{}
                err    = json.Unmarshal(value, &dynima)
                if err == nil {
                    if( dynima.SourcePath == path) {
                        fmt.Printf("\nMatched %s\n",path)
                        dynimas = append(dynimas,dynima)
                    }
                }
                return nil
            })
            if len(dynimas) == 0 {
                return collection_entry_list_is_empty
            }
            return err
        });
        if err == nil {  return dynimas, err } else {  return nil,err }
    }
    return nil, err
}

