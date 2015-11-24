package sistory
import "github.com/boltdb/bolt"

var db_path string =  "/tmp/sis.db"

type Storage struct {
    Db    *bolt.DB
    Error bool
}

func Open()(s *Storage) {
    db, err := bolt.Open(db_path, 0600, nil)
    if err!= nil { s.Error = true ; return } else { s.Db = db }
    return

}

func (s *Storage) Close () {

    s.Db.Close()

}
