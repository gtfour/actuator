package userstorage

import "wapour/settings"
import "wapour/api/webclient"
import "github.com/boltdb/bolt"

var UsersStorage webclient.WengineWrapperStorage

type UserStorage interface {

    FindWrapper(err error)
    AddWrapper()



}


type FileUserStorage struct {
    StorageType      string
    Wrappers         *[]*webclient.WengineWrapper
}

type BoltdbUserStorage   struct {
   StorageType      string
   StorageFileName  string
   Db bolt.DB
}

func init() {
    var db_open_failed bool
    if settings.ONLINE_USERS_STORAGE_TYPE        == "db" {
        UsersStorage.StorageType = "db"
        db_path                  := settings.ONLINE_USERS_DB_FILE
        db, err                  := bolt.Open(db_path, 0600, nil)
        if err!= nil {
            defer db.Close()
            db_open_failed = true
        }
    } else if settings.ONLINE_USERS_STORAGE_TYPE == "ram" || db_open_failed == true {
        UsersStorage.StorageType = "ram"
    }
}

func main() {

}

