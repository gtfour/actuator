package sistory
import "github.com/boltdb/bolt"
import "fmt"

var db_path string =  "/tmp/sis.db"
var comments =  []string {`//` , `#`}

type Storage struct {
    Db    *bolt.DB
    Error bool
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


func Open()(s Storage) {
    db, err := bolt.Open(db_path, 0600, nil)
    if err!= nil { s.Error = true ; return } else { s.Db = db }
    return

}

func (s *Storage) Close () {

    s.Db.Close()

}

func(s *Storage) CallSpirit (path string) (data []byte)  {


    
    s.Db.View(func (tx *bolt.Tx) error {
        bucket:=tx.Bucket([]byte(path))
        if bucket == nil { fmt.Printf("Bucket is nil") ; return nil}
        data=bucket.Get([]byte("hello_key"))
        return nil
    })

    return data

}

func CreateNewbie (path string)(sp SpiritProp)  {


    return



}

func Compare( newbie, spirit *SpiritProp ) (difference []string)  {


    return



}

func(s *Storage) UploadSpirit (sp *SpiritProp) (err error) {

    s.Db.Update( func(tx *bolt.Tx) error {
        // replace existing bucket if exists
        bucket:=tx.Bucket([]byte(sp.Path))
        if bucket != nil { /* if bucket exists - remove it */
            err = tx.DeleteBucket([]byte(sp.Path))
            if err != nil { return err }
        }
        bucket, err =tx.CreateBucket([]byte(sp.Path))
        if err!= nil { return err }
        return bucket.Put([]byte("hello_key"),[]byte("hello_value"))
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
