package sistory
import "github.com/boltdb/bolt"
//import "fmt"

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
    Lines          []string
    IgnoreComment  bool

}

type Difference struct {

    




}


func Open()(s Storage) {
    db, err := bolt.Open(db_path, 0600, nil)
    if err!= nil { s.Error = true ; return } else { s.Db = db }
    return

}

func (s *Storage) Close () {

    s.Db.Close()

}

func(s *Storage) CallSpirit(path string) {


    return



}

func CreateNewbie (path string)(sp SpiritProp)  {


    return



}

func Compare( newbie, spirit *SpiritProp ) (difference []string)  {


    return



}

func (sp *SpiritProp)  UploadSpirit() {



}
