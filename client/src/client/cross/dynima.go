package cross

import "fmt"
import "errors"
import "encoding/json"
import "github.com/boltdb/bolt"

var dynima_edit_error =  errors.New("Unable to edit dynima")
var dynima_get_error  =  errors.New("Unable to get dynima")

type Dynima struct {
    //parsers
    Id              string
    ids             []int // id column number
    SourcePath      string
    SourceType      string
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

    //fmt.Printf("Storage error %v\nNew dynima id %s\n",STORAGE_INSTANCE.Error,id)
    return nil

}

func (d *Dynima) Write()(err error){
    if STORAGE_INSTANCE.Error == false {

        db:=STORAGE_INSTANCE.Db
        err=db.Update(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(STORAGE_INSTANCE.dynimasTableName))
            if b==nil{ return nil }
            encoded, err := json.Marshal(d)
            if err!=nil{ return err }
            return b.Put([]byte(d.Id), encoded) //CreateBucket has been replaced to CreateBucketIfNotExists because when err==bolt.ErrBucketExists - dynima is nil
         //// if err==nil || err==bolt.ErrBucketExists { // If the key exist then its previous value will be overwritten
                //fmt.Printf("\nEdit dynima:\n%v\nError: %v\n",d,err)
                ////err=dynima.Put([]byte("source_path"),[]byte(d.SourcePath))
                ////if err!=nil{ return err }
                ////err=dynima.Put([]byte("source_type"),[]byte(d.SourceType))
                ////if err!=nil{ return err }
                ////err=dynima.Put([]byte("template"),[]byte(d.template))
                ////if err!=nil{ return err }
             //// encoded, err := json.Marshal(d)
            ////  if err != nil {
            ////      return err
           ////   }
          ////    return dynima.Put([]byte(user.Name), encoded)
            ////} else { return err }
            return nil
        });

    }
    return dynima_edit_error
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
            if b==nil{ return nil }
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
    return dynima_edit_error
}

func RemoveDynima(id string)(error){
    return nil
}


func GetDynima(id string)(*Dynima,error){
    var err error
    if STORAGE_INSTANCE.Error == false {

        db     := STORAGE_INSTANCE.Db
        dynima := &Dynima{}
        err = db.View(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(STORAGE_INSTANCE.dynimasTableName))
            //fmt.Printf("\nDynimas collection doesnt exist\n")
            if b==nil{ return dynima_get_error }
            data:=b.Get([]byte(id))
            err = json.Unmarshal(data, &dynima)
            return err
            //fmt.Printf("\nDynima doesnt exist\n")
            ////if d==nil{ return dynima_get_error }
            ////dynima = &Dynima{}
            ////source_path := d.Get([]byte("source_path"))
            ////if source_path == nil { source_path=[]byte("")  }
            ////source_type      := d.Get([]byte("source_type"))
            ////if source_type == nil { source_type=[]byte("")  }
            ////template     := d.Get([]byte("template"))
            ////if template == nil { template=[]byte("")  }
            ////dynima.SourcePath   = id
            ////dynima.SourcePath   = string(source_path)
            ////dynima.SourceType   = string(source_type)
            ////dynima.template     = string(template)
        });
        if err == nil { return dynima, nil } else { return nil, dynima_get_error }
    }
    return nil, dynima_get_error
}

func GetDynimasByPath(path string) ( dynimas  []Dynima , err  error ) {

    if STORAGE_INSTANCE.Error == false {

        db     := STORAGE_INSTANCE.Db
        dynimas = make([]Dynima,0)
        //dynima := &Dynima{}
        err = db.View(func(tx *bolt.Tx) error {
            b:=tx.Bucket([]byte(STORAGE_INSTANCE.dynimasTableName))
            //fmt.Printf("\nDynimas collection doesnt exist\n")
            if b==nil{ return dynima_get_error }
            //data:=b.Get([]byte(id))
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




            //err = json.Unmarshal(data, &dynima)
            return err
            //fmt.Printf("\nDynima doesnt exist\n")
            ////if d==nil{ return dynima_get_error }
            ////dynima = &Dynima{}
            ////source_path := d.Get([]byte("source_path"))
            ////if source_path == nil { source_path=[]byte("")  }
            ////source_type      := d.Get([]byte("source_type"))
            ////if source_type == nil { source_type=[]byte("")  }
            ////template     := d.Get([]byte("template"))
            ////if template == nil { template=[]byte("")  }
            ////dynima.SourcePath   = id
            ////dynima.SourcePath   = string(source_path)
            ////dynima.SourceType   = string(source_type)
            ////dynima.template     = string(template)
        });
        if err == nil { return dynimas, err } else { return nil,err }
    }
    return nil, err
}

