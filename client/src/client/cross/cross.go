package cross

import "fmt"
import "io/ioutil"
import "github.com/boltdb/bolt"
import "client/settings"

var comments =  []string {`//` , `#`}

var Database = Open()



type Storage struct {
    Db                 *bolt.DB
    Error              error
    dynimasTableName   string
    dashgatesTableName string
    motionsTableName   string
}



type SpiritProp struct {

    Path           string
    Type           string
    Seek           uint64 // just for log-files
    Size           uint64
    Lines          []string
    IgnoreComment  bool

}

type Difference struct {
    field string
}


func Open()(s Storage){
    fmt.Printf("\nTrying to open cross database: %s\n",settings.SYSTEM_DATABASE)
    db, err := bolt.Open(settings.SYSTEM_DATABASE, 0600, nil)
    if err!= nil { s.Error = err ; return } else {
        s.Db                = db
        s.dynimasTableName   = "dynimas"
        s.dashgatesTableName = "dashgates"
        s.motionsTableName   = "motions"
        // init collections
        db.Update(func(tx *bolt.Tx) error {
            //_,err_dynimas    := tx.CreateBucket([]byte(s.dynimasTableName))
            //_,err_dashgates :=  tx.CreateBucket([]byte(s.dashgatesTableName))
            _,err_dynimas    := tx.CreateBucketIfNotExists([]byte(s.dynimasTableName))
            _,err_dashgates  := tx.CreateBucketIfNotExists([]byte(s.dashgatesTableName))
            _,err_motions    := tx.CreateBucketIfNotExists([]byte(s.motionsTableName))
            if err_dynimas != nil || err_dashgates != nil || err_motions != nil {
                s.Error = unable_to_init
                defer db.Close()
                return unable_to_init
            }
            return nil
        });
        s.Error             = nil
    }
    return
}

func (s *Storage) Close () {
    s.Db.Close()
}

func(s *Storage) CallSpirit (path string) (data []byte)  {


    s.Db.View(func (tx *bolt.Tx) error {
        bucket:=tx.Bucket([]byte(path))
        if bucket == nil { fmt.Printf("Bucket is nil") ; return nil }
        data=bucket.Get([]byte("content"))
        return nil
    })

    return data

}

func CreateNewbie (path string)(sp SpiritProp)  {

    content, err := ioutil.ReadFile(sp.Path)
    if err!= nil { return sp }
    fmt.Printf("%s",content)



    return sp



}

func Compare( newbie, spirit *SpiritProp ) (difference []string)  {

    return



}

func(s *Storage) UploadSpirit (sp *SpiritProp) (err error) {

    content, err := ioutil.ReadFile(sp.Path)
    fmt.Printf("%s",content)
    if err!= nil { return err }

    s.Db.Update( func(tx *bolt.Tx) error {
        // replace existing bucket if exists
        bucket:=tx.Bucket([]byte(sp.Path))
        if bucket != nil { /* if bucket exists - remove it */
            err = tx.DeleteBucket([]byte(sp.Path))
            if err != nil { return err }
        }
        bucket, err =tx.CreateBucket([]byte(sp.Path))
        if err!= nil { return err }
        return bucket.Put([]byte("content"),[]byte(`{"192.168.236.11":"controller","127.0.0.1":"localhost"}`))
    })

    /*s.db.Update(func(tx *bolt.Tx) error {
    b, err := tx.CreateBucketIfNotExists([]byte(sp.Path))
    if err != nil {
        return err
    }
    return b.Put([]byte("2015-01-01"), []byte("My New Year post"))
    })*/

    return nil

}
